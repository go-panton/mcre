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
	Create(string) error
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

	//1st get file name
	fi, err := fp.Stat()
	if err != nil {
		return err
	}

	fileSize := fi.Size()
	startDate := fi.ModTime()
	fileDT := fi.ModTime().Format("2006-01-02") //convert to string and designated format

	//Get extension of fileServer
	fileExt := path.Ext(filename)

	newNode, err := nodes.NewNode(nodeID, filename, fileDT)
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

	// use defer instead of manually closing it due to cases where service failed halfway
	defer destFile.Close()

	if err != nil {
		return err
	}

	_, err = io.Copy(destFile, fp)
	if err != nil {
		return err
	}

	//Create New Fmedia record
	newFmedia, err := fmedias.NewFmedia(nodeID, filename, newFileName, fileExt, storageFolder, filename, int(fileSize))
	if err != nil {
		return err
	}

	//Create Nodelink
	//Hardcoded the folder nodeID to 200004 and link type to FILE
	newNodelink, err := nodes.NewNodelink(nodeID, folderNodeID, linkType)
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

	//Create Fverinfo, treat it as first version since no possible way to know what version is that file
	version := "v1"
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

	err = svc.fmediaR.CreateTx(nodeStr, fmediaStr, nlStr, fverStr, convStr, fmpStrSlice)
	if err != nil {
		//Remove the file from storage as the transaction has already failed
		err := os.Remove(dest)
		return err
	}

	return nil
}
