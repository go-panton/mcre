package files

import (
	"errors"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	fmedias "github.com/go-panton/mcre/fmedias/models"
	nodes "github.com/go-panton/mcre/nodes/models"
	seqs "github.com/go-panton/mcre/seqs/models"
)

// BadRequestError is a http-error-wrapper of 404
type BadRequestError error

// Service is the interface of download API
type Service interface {
	//Download(simple) opens and reads file based on fileid
	Download(fileid string) io.Reader
	//Create copy the file to server and store file's metadata onto database
	Create(string) error
	//Version copy the file to server and create a new version based on existing file
	Version(string, int) error
}

// NewService instantiates new download-service.
func NewService(seqR seqs.SeqRepository, nodeR nodes.NodeRepository, fmediaR fmedias.FmediaRepository, nodelinkR nodes.NodelinkRepository, fverinfoR fmedias.FverinfoRepository, convqueueR fmedias.ConvqueueRepository, fmpolicyR fmedias.FmpolicyRepository) Service {
	return &service{seqR, nodeR, fmediaR, nodelinkR, fverinfoR, convqueueR, fmpolicyR}
}

type service struct {
	seqR       seqs.SeqRepository
	nodeR      nodes.NodeRepository
	fmediaR    fmedias.FmediaRepository
	nodelinkR  nodes.NodelinkRepository
	fverinfoR  fmedias.FverinfoRepository
	convqueueR fmedias.ConvqueueRepository
	fmpolicyR  fmedias.FmpolicyRepository
}

func (svc *service) Download(fileid string) io.Reader {
	return strings.NewReader("This is your requested file content.")
}

func (svc *service) Create(filename string) error {
	if filename == "" {
		return errors.New("Filename passing in is empty!")
	}

	nodeID, err := svc.seqR.Find("NODE")
	if err != nil {
		return err
	}
	fileID, err := svc.seqR.Find("FILENAME")
	if err != nil {
		return err
	}

	//open the file to get the file pointer
	fp, err := os.Open(filename)

	// although it is not the best to use defer for close because the file can be close much earlier
	// but we have to consider when error happen we want to make sure the close function always get called
	defer fp.Close()

	//1st get file's metadata
	fi, err := fp.Stat()
	if err != nil {
		return err
	}

	//fname is file's name without path before it
	fname := fi.Name()
	fileSize := fi.Size()
	startDate := fi.ModTime()
	fileDT := fi.ModTime().Format("2006-01-02") //convert to string and designated format

	//Get extension of fileServer
	fileExt := path.Ext(filename)

	newNode, err := nodes.NewNode(nodeID, fname, fileDT)
	if err != nil {
		return errors.New("Node creation error: " + err.Error())
	}

	//Copy File to designated storage folder
	const storageFolder = "D:\\PantonSys\\PTD\\ISOStorage\\1\\"

	const folderNodeID = 200004

	const linkType = "FILE"

	newFileName := strconv.Itoa(fileID) + fileExt

	dest := storageFolder + newFileName

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	// use defer instead of manually closing it due to cases where service failed halfway
	defer destFile.Close()

	defer os.Remove(dest)

	_, err = io.Copy(destFile, fp)
	if err != nil {
		return err
	}

	//Create New Fmedia record
	newFmedia, err := fmedias.NewFmedia(nodeID, fname, newFileName, fileExt, storageFolder, fname, int(fileSize))
	if err != nil {
		return err
	}

	//Create Nodelink
	//Hardcoded the folder nodeID to 200004 and link type to FILE
	newNodelink, err := nodes.NewNodelink(nodeID, folderNodeID, linkType)
	if err != nil {
		return err
	}

	//Create Fverinfo, treat it as first version since no possible way to know what version is that file
	version := "1"
	endDate := startDate.Add(168 * time.Hour).Format("2006-01-02")
	verState := 1
	remarks := "Testing Version 1"

	newFverinfo, err := fmedias.NewFverinfo(nodeID, fileDT, endDate, version, verState, remarks)
	if err != nil {
		return err
	}

	//Create Convqueue
	newConvqueue, err := fmedias.NewConvqueue(nodeID, time.Now().Format("2006-01-02 15:04:05"), fileExt, dest)
	if err != nil {
		return err
	}

	//Find folder's fmpolicy 1st
	folderSec, err := svc.fmpolicyR.FindUsingNodeID(folderNodeID)

	//slice of fmp insert string
	var fmpStrSlice []string

	//Create new fmpolicy for every instance
	for _, fs := range folderSec {
		newFmp, e := fmedias.NewFmpolicy(fs.FmpDownload, fs.FmpRevise, fs.FmpView, fs.FmpUGID, fs.FmpUGType, nodeID)
		if e != nil {
			return e
		}
		fmpStr, e := svc.fmpolicyR.GetInsertStr(newFmp)
		if e != nil {
			return e
		}
		fmpStrSlice = append(fmpStrSlice, fmpStr)
	}
	//meaning folder has no fmpolicy
	if len(fmpStrSlice) == 0 {
		newFmp, e := fmedias.NewFmpolicy(1, 1, 1, 0, 0, nodeID)
		if e != nil {
			return e
		}

		fmpStr, e := svc.fmpolicyR.GetInsertStr(newFmp)
		if e != nil {
			return e
		}
		fmpStrSlice = append(fmpStrSlice, fmpStr)
	}

	// Gather all insert statement for transaction
	nodeStr, err := svc.nodeR.GetInsertStr(newNode)
	if err != nil {
		return err
	}
	fmediaStr, err := svc.fmediaR.GetInsertStr(newFmedia)
	if err != nil {
		return err
	}
	nlStr, err := svc.nodelinkR.GetInsertStr(newNodelink)
	if err != nil {
		return err
	}
	fverStr, err := svc.fverinfoR.GetInsertStr(newFverinfo)
	if err != nil {
		return err
	}
	convStr, err := svc.convqueueR.GetInsertStr(newConvqueue)
	if err != nil {
		return err
	}

	if err = svc.fmediaR.CreateFileTx(nodeStr, fmediaStr, nlStr, fverStr, convStr, fmpStrSlice); err != nil {
		//Remove the file from storage as the transaction has already failed
		return err
	}

	if err = svc.seqR.Update("NODE", nodeID); err != nil {
		return err
	}

	if err = svc.seqR.Update("FILENAME", fileID); err != nil {
		return err
	}

	return nil
}

func (svc *service) Version(filename string, currentFileNID int) error {
	if filename == "" {
		return errors.New("Filename passing in is empty!")
	}

	nodeID, err := svc.seqR.Find("NODE")
	if err != nil {
		return err
	}
	fileID, err := svc.seqR.Find("FILENAME")
	if err != nil {
		return err
	}

	//open the file to get the file pointer
	fp, err := os.Open(filename)

	defer fp.Close()

	fi, err := fp.Stat()
	if err != nil {
		return err
	}

	fname := fi.Name()
	fileSize := fi.Size()
	startDate := fi.ModTime()
	fileDT := fi.ModTime().Format("2006-01-02") //convert to string and designated format

	//Get extension of fileServer
	fileExt := path.Ext(filename)

	newNode, err := nodes.NewNode(nodeID, fname, fileDT)
	if err != nil {
		return errors.New("Node creation error: " + err.Error())
	}

	const storageFolder = "D:\\PantonSys\\PTD\\ISOStorage\\1\\"

	const folderNodeID = 200004

	const linkType = "FILE"

	newFileName := strconv.Itoa(fileID) + fileExt

	dest := storageFolder + newFileName

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destFile.Close()

	defer os.Remove(dest)

	_, err = io.Copy(destFile, fp)
	if err != nil {
		return err
	}

	// Retrieve current file's data from database
	currentFmedia, err := svc.fmediaR.Find(currentFileNID)
	if err != nil {
		return err
	}

	currentFver, err := svc.fverinfoR.Find(currentFileNID)
	if err != nil {
		return err
	}

	// List all version of current file using node ID and link type
	childList, err := svc.nodelinkR.FindByParent(currentFileNID, "VER_LINKS")

	//Create New Fmedia record
	newFmedia, err := fmedias.NewFmedia(nodeID, fname, newFileName, fileExt, storageFolder, fname, int(fileSize))
	if err != nil {
		return err
	}

	newNodelink, err := nodes.NewNodelink(nodeID, folderNodeID, linkType)
	if err != nil {
		return err
	}

	// Since the database column specify version to be varchar we have to convert it back to int to increment the version number
	// If we decide to change the value of column to int then we will be able to just increment the version number
	version, err := strconv.Atoi(currentFver.Version)
	if err != nil {
		return err
	}
	version++

	verString := strconv.Itoa(version)
	if err != nil {
		return err
	}
	endDate := startDate.Add(168 * time.Hour).Format("2006-01-02") //as of now only add 1 week
	verState := 1
	remarks := "Testing Version 2"

	newFverinfo, err := fmedias.NewFverinfo(nodeID, fileDT, endDate, verString, verState, remarks)
	if err != nil {
		return err
	}

	var deleteNLSlice []string
	var insertNLSlice []string

	if len(childList) != 0 {
		// Unlink them and relink to new file
		for _, childNID := range childList {
			deleteStr, e := svc.nodelinkR.GetDeleteStr(childNID, currentFileNID)
			if e != nil {
				return e
			}

			deleteNLSlice = append(deleteNLSlice, deleteStr)

			newnl, e := nodes.NewNodelink(childNID, nodeID, "VER_LINKS")
			if e != nil {
				return e
			}
			insertStr, e := svc.nodelinkR.GetInsertStr(newnl)
			if e != nil {
				return e
			}
			insertNLSlice = append(insertNLSlice, insertStr)
		}
	}

	// Current file is now a version of new file
	newnl, err := nodes.NewNodelink(currentFileNID, nodeID, "VER_LINKS")
	if err != nil {
		return err
	}

	insertNlStr, err := svc.nodelinkR.GetInsertStr(newnl)
	if err != nil {
		return err
	}

	insertNLSlice = append(insertNLSlice, insertNlStr)

	// Unlink current file with its parent folder, this is assuming new version will be at the same folder with current version
	// , if not change folder nodeID to where the current version's folder node ID
	deleteNLStr, err := svc.nodelinkR.GetDeleteStr(currentFileNID, folderNodeID)
	if err != nil {
		return err
	}

	deleteNLSlice = append(deleteNLSlice, deleteNLStr)

	// Note to anyone seeing this: The process of unlink and relink CAN be simplified into 1 update query BUT
	// the legacy code I refer to, remove and create a new nodelink the saying goes: "if it ain't break don't fix it!"

	//Update Status of File to 0. I have no idea what 0 means since legacy code does not have 0 in FileStatus. somone screwed up and
	//never initialized the fstatus value hence it becomes 0

	//TODO: Revisit Fmedia Fstatus and Ftype, currently they are meaningless to be included, would need to discuss whether to rmeove it or retain it
	currentFmedia.FStatus = 0

	updateCurrFmStr, err := svc.fmediaR.GetUpdateStr(currentFmedia)
	if err != nil {
		return err
	}

	//Verstate 2 means inactive
	currentFver.VerState = 2

	if currentFver.VerState == 3 {
		currentFver.EndDate = time.Now().Format("2006-01-02")
	} else if currentFver.VerState == 1 {
		currentFver.StartDate = time.Now().Format("2006-01-02")
		currentFver.EndDate = ""
	}

	updateCurrFvStr, err := svc.fverinfoR.GetUpdateStr(currentFver)
	if err != nil {
		return err
	}

	fileSec, err := svc.fmpolicyR.FindUsingNodeID(currentFileNID)
	if err != nil {
		return err
	}

	var fmpSlice []string

	for _, s := range fileSec {
		newFmp, e := fmedias.NewFmpolicy(s.FmpDownload, s.FmpRevise, s.FmpView, s.FmpUGID, s.FmpUGType, nodeID)
		if e != nil {
			return e
		}

		fmpinsertStr, e := svc.fmpolicyR.GetInsertStr(newFmp)
		if e != nil {
			return e
		}
		fmpSlice = append(fmpSlice, fmpinsertStr)
	}

	//Create Convqueue
	newConvqueue, err := fmedias.NewConvqueue(nodeID, time.Now().Format("2006-01-02 15:04:05"), fileExt, dest)
	if err != nil {
		return err
	}

	nodeStr, err := svc.nodeR.GetInsertStr(newNode)
	if err != nil {
		return err
	}
	fmStr, err := svc.fmediaR.GetInsertStr(newFmedia)
	if err != nil {
		return err
	}
	nlStr, err := svc.nodelinkR.GetInsertStr(newNodelink)
	if err != nil {
		return err
	}
	fvStr, err := svc.fverinfoR.GetInsertStr(newFverinfo)
	if err != nil {
		return err
	}
	convqStr, err := svc.convqueueR.GetInsertStr(newConvqueue)
	if err != nil {
		return err
	}

	dbStmt := fmedias.Statement{
		CreateNodeStmt:   nodeStr,
		CreateFmStmt:     fmStr,
		CreateNlStmt:     nlStr,
		CreateFvStmt:     fvStr,
		CreateConvStmt:   convqStr,
		CreateNLSlice:    insertNLSlice,
		DeleteNLSlice:    deleteNLSlice,
		UpdateCurrFmStmt: updateCurrFmStr,
		UpdateCurrFvStmt: updateCurrFvStr,
		CreateFmpSlice:   fmpSlice,
	}

	err = svc.fmediaR.CreateVersionTx(dbStmt)
	if err != nil {
		return err
	}

	if err = svc.seqR.Update("NODE", nodeID); err != nil {
		return err
	}

	if err = svc.seqR.Update("FILENAME", fileID); err != nil {
		return err
	}

	return nil
}
