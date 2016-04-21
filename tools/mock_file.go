package tools

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Node struct {
	NodeID    int
	NodeBits  int
	NodeDesc  string
	NodeDT    string //t.format("2006-01-02")
	NodeGID   int
	NodeHash  string
	NodeLevel int
	NodeType  string
	NodeUID   int
}

type Fmedia struct {
	NodeID   int
	FDesc    string
	FExt     string
	FFulPath string
	FGName   string
	FOName   string
	FRemark  string
	FSize    int
	FStatus  int
	FType    int
}

type Fverinfo struct {
	NodeID    int
	EndDate   string //t.format("2006-01-02")
	Remarks   string
	StartDate string //t.format("2006-01-02")
	Version   string
	VerState  int
}

type Convqueue struct {
	NodeID   int
	ConvType string
	FExt     string
	FFulPath string
	INSDate  string ////t.format("2006-01-02 12:47:22")
	Priority int
}

type NodeLink struct {
	LinkCNodeID int
	LinkPNodeID int
	LinkType    string
}

//mock_insert_file mock a file insert operation in mcre
func mock_insert_file() error {
	const (
		FILE_NAME = "test1.txt"
		FILE_EXT  = ".txt"

		FILE_SIZE = 1024

		STORAGE_FOLDER = "D:\\PantonSys\\PTD\\ISOStorage\\1\\"

		DATE_ONLY = "2006-01-02"

		DATE_AND_TIME = "2006-01-02 15:04:05"

		GRAPH_ID = 200001

		FILE_TYPE = "FILE"

		USER_ID = 1

		VER_STATE = 1
	)

	fileID, err := IDGenerator("FILENAME")
	if err != nil {
		return err
	}
	//Increment by 1 indicate a new file created
	fileID += 1

	fmt.Println("New FILE ID: ", fileID)

	var TIME_START string = time.Now().Format(DATE_ONLY)
	var TIME_AFTER string = time.Now().Add(150 * time.Hour).Format(DATE_ONLY)
	var FG_NAME string = strconv.Itoa(fileID) + FILE_EXT
	var DEST_PATH string = STORAGE_FOLDER + FG_NAME

	nodeID, err := IDGenerator("NODE")
	if err != nil {
		return err
	}
	//Increment by 1 indicate a new node created
	nodeID++

	fmt.Println("New Node ID: ", nodeID)

	//create a new file in order to be copy over
	newFile, err := os.Create(FILE_NAME)
	if err != nil {
		return err
	}
	newFile.Close()

	//truncate the file to FILE_SIZE
	err1 := os.Truncate(FILE_NAME, FILE_SIZE)
	if err1 != nil {
		return err1
	}

	//now to fake data into struct
	newNode := Node{
		NodeID:    nodeID,
		NodeBits:  0,
		NodeDesc:  FILE_NAME,
		NodeDT:    TIME_START,
		NodeGID:   GRAPH_ID,
		NodeHash:  "",
		NodeLevel: -32768,
		NodeType:  FILE_TYPE,
		NodeUID:   USER_ID,
	}

	newFmedia := Fmedia{
		NodeID:   nodeID,
		FDesc:    FILE_NAME,
		FExt:     FILE_EXT,
		FFulPath: STORAGE_FOLDER,
		FGName:   FG_NAME,
		FOName:   FILE_NAME,
		FRemark:  "",
		FSize:    FILE_SIZE,
		FStatus:  1,
		FType:    1,
	}

	newFverInfo := Fverinfo{
		NodeID:    nodeID,
		EndDate:   TIME_AFTER,
		Remarks:   "This is test version",
		StartDate: TIME_START,
		Version:   "1",
		VerState:  VER_STATE,
	}

	newConQ := Convqueue{
		NodeID:   nodeID,
		ConvType: "PDF",
		FExt:     FILE_EXT,
		FFulPath: DEST_PATH,
		INSDate:  time.Now().Format(DATE_AND_TIME),
		Priority: 1,
	}

	newNL := NodeLink{
		LinkPNodeID: 200004,
		LinkCNodeID: nodeID,
		LinkType:    "FILE",
	}

	//Now to copy the file to destination path
	originalFile, err := os.Open(FILE_NAME)
	if err != nil {
		return errors.New("Unable to open file")
	}

	destFile, err := os.Create(DEST_PATH)
	if err != nil {
		return errors.New("unable to create a new file")
	}

	bytesWritten, err := io.Copy(destFile, originalFile)
	if err != nil {
		return errors.New("Unable to copy file")
	}
	fmt.Println("Bytes copied: ", bytesWritten)

	originalFile.Close()
	destFile.Close()

	//Now to insert the struct into database
	db, err := sql.Open("mysql", "root:root123@/ptd_new")
	if err != nil {
		fmt.Errorf("Error creating connection towards database")
		return err
	}

	stmtN, err := db.Prepare("INSERT node SET nodeid=?,nodebits=?,nodedesc=?,nodedt=?,nodegid=?,nodehash=?,nodelevel=?,nodetype=?,nodeuid=?")
	if err != nil {
		fmt.Errorf("Error creating insert statement for node")
		return err
	}
	stmtFmedia, err := db.Prepare("INSERT fmedia SET nodeid=?,fdesc=?,fext=?,ffulpath=?,fgname=?,foname=?,fremark=?,fsize=?,fstatus=?,ftype=?")
	if err != nil {
		fmt.Errorf("Error creating insert statement for Fmedia")
		return err
	}
	stmtFverinfo, err := db.Prepare("INSERT fverinfo SET nodeid=?,enddate=?,remarks=?,startdate=?,version=?,verstate=?")
	if err != nil {
		fmt.Errorf("Error creating insert statement for fverinfo")
		return err
	}
	stmtconvqueue, err := db.Prepare("INSERT convqueue SET nodeid=?,convtype=?,fext=?,ffulpath=?,insdate=?,priority=?")
	if err != nil {
		fmt.Errorf("Error creating insert statement for convqueue")
		return err
	}
	stmtnodelink, err := db.Prepare("INSERT nodelink SET linkcnodeid=?,linkpnodeid=?,linktype=?")
	if err != nil {
		fmt.Errorf("Error creating insert statement for nodelink")
		return err
	}

	resN, err := stmtN.Exec(newNode.NodeID, newNode.NodeBits, newNode.NodeDesc, newNode.NodeDT, newNode.NodeGID, newNode.NodeHash, newNode.NodeLevel, newNode.NodeType, newNode.NodeUID)
	if err != nil {
		fmt.Errorf("Error inserting data into Node table")
		return err
	}
	fmt.Println("Result: ", resN)

	resFm, err := stmtFmedia.Exec(newFmedia.NodeID, newFmedia.FDesc, newFmedia.FExt, newFmedia.FFulPath, newFmedia.FGName, newFmedia.FOName, newFmedia.FRemark, newFmedia.FSize, newFmedia.FStatus, newFmedia.FType)
	if err != nil {
		fmt.Errorf("Error inserting data into Fmedia table")
		return err
	}
	fmt.Println("Result: ", resFm)

	resFv, err := stmtFverinfo.Exec(newFverInfo.NodeID, newFverInfo.EndDate, newFverInfo.Remarks, newFverInfo.StartDate, newFverInfo.Version, newFverInfo.VerState)
	if err != nil {
		fmt.Errorf("Error inserting data into Fverinfo table")
		return err
	}
	fmt.Println("Result: ", resFv)

	resCv, err := stmtconvqueue.Exec(newConQ.NodeID, newConQ.ConvType, newConQ.FExt, newConQ.FFulPath, newConQ.INSDate, newConQ.Priority)
	if err != nil {
		fmt.Errorf("Error inserting data into convqueue table")
		return err
	}
	fmt.Println("Result: ", resCv)

	resnl, err := stmtnodelink.Exec(newNL.LinkCNodeID, newNL.LinkPNodeID, newNL.LinkType)
	if err != nil {
		fmt.Errorf("Error inserting data into nodelink table")
		return err
	}
	fmt.Println("Result: ", resnl)

	db.Close()

	return nil
}
