package files

import (
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
	Create(file *os.File) error
}

// NewService instantiates new download-service.
func NewService(seq seqs.SeqRepository, node nodes.NodeRepository, fmedia fmedias.FmediaRepository, nodelink nodes.NodelinkRepository, fverinfo fmedias.FverinfoRepository, convqueue fmedias.ConvqueueRepository) Service {
	return &service{seq, node, fmedia, nodelink, fverinfo, convqueue}
}

type service struct {
	seq       seqs.SeqRepository
	node      nodes.NodeRepository
	fmedia    fmedias.FmediaRepository
	nodelink  nodes.NodelinkRepository
	fverinfo  fmedias.FverinfoRepository
	convqueue fmedias.ConvqueueRepository
}

func (svc *service) Download(fileid string) io.Reader {
	return strings.NewReader("This is your requested file content.")
}

func (svc *service) Create(fp *os.File) error {

	nodeID, err := svc.seq.Find("NODEID")
	if err != nil {
		return err
	}
	fileID, err := svc.seq.Find("FILENAME")
	if err != nil {
		return err
	}

	//1st get file name
	fi, err := fp.Stat()
	if err != nil {
		return err
	}

	fileName := fi.Name()
	fileSize := fi.Size()
	startDate := fi.ModTime()
	fileDT := fi.ModTime().Format("2006-01-02") //convert to string and designated format

	//Get extension of fileServer
	fileExt := path.Ext(fileName)

	newNode, err := nodes.NewNode(nodeID, fileName, fileDT)
	if err != nil {
		return err
	}

	//Copy File to designated storage folder
	const storageFolder = "D:\\PantonSys\\PTD\\ISOStorage\\1\\"

	newFileName := strconv.Itoa(fileID) + fileExt

	dest := storageFolder + newFileName

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	_, err = io.Copy(destFile, fp)
	if err != nil {
		return err
	}
	destFile.Close()

	//Create New Fmedia record
	newFmedia, err := fmedias.NewFmedia(nodeID, fileName, newFileName, fileExt, storageFolder, fileName, int(fileSize))
	if err != nil {
		return err
	}

	//Create Nodelink
	newNodelink, err := nodes.NewNodelink(nodeID)
	if err != nil {
		return err
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

	//Should I make a trasaction instead....
	nodeStr, err := svc.node.GetInsertStr(newNode)
	if err != nil {
		return err
	}
	fmediaStr, err := svc.fmedia.GetInsertStr(newFmedia)
	if err != nil {
		return err
	}
	nlStr, err := svc.nodelink.GetInsertStr(newNodelink)
	if err != nil {
		return err
	}
	fverStr, err := svc.fverinfo.GetInsertStr(newFverinfo)
	if err != nil {
		return err
	}
	convStr, err := svc.convqueue.GetInsertStr(newConvqueue)
	if err != nil {
		return err
	}

	err = svc.fmedia.CreateTx(nodeStr, fmediaStr, nlStr, fverStr, convStr)
	if err != nil {
		//Remove the file from storage as the transaction has already failed
		err := os.Remove(dest)
		return err
	}

	return nil
}
