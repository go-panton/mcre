package tools

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"time"

	"strconv"

	"errors"

	"github.com/go-panton/mcre/infra/store/mysql"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Node struct define the data in the node table
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

//Fmedia define the data inside Fmedia table
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

//Fverinfo struct define the data in the Fverinfo table
type Fverinfo struct {
	NodeID    int
	EndDate   string //t.format("2006-01-02")
	Remarks   string
	StartDate string //t.format("2006-01-02")
	Version   string
	VerState  int
}

//Convqueue struct define the data inside Convqueue table
type Convqueue struct {
	NodeID   int
	ConvType string
	FExt     string
	FFulPath string
	INSDate  string //t.format("2006-01-02 12:47:22")
	Priority int
}

//NodeLink struct define the data in nodelink table
type NodeLink struct {
	LinkCNodeID int
	LinkPNodeID int
	LinkType    string
}

//PreparedStmt struct store all the sql statement that will be execute later on
type PreparedStmt struct {
	NodeStmt   *sql.Stmt
	FmediaStmt *sql.Stmt
	NLStmt     *sql.Stmt
	FvStmt     *sql.Stmt
	ConvStmt   *sql.Stmt
}

const (
	//dbUserName is the username for the database
	dbUserName = "root"

	//dbPassword is the password for the database
	dbPassword = "root123"

	//dbName is the name of the database
	dbName = "ptd_new"

	//fileName is the name of the file that is going to be created
	fileName = "test2.txt"

	//fileExt is the extension of the file
	fileExt = ".txt"

	//fileSize is the size of the file
	fileSize = 2048

	//storageFolder define the folder where all files get stored in
	storageFolder = "D:\\PantonSys\\PTD\\ISOStorage\\1\\"

	//dateOnly format the time to use date only
	dateOnly = "2006-01-02"

	//dateAndTime format the time to have date and time only
	dateAndTime = "2006-01-02 15:04:05"

	//graphID is the graph for the file
	graphID = 200001

	//fileType is the file type
	fileType = "FILE"

	//userID is the user id for the file created
	userID = 1

	//verState is the version state of the file
	verState = 1

	//parentID is the parent folder node id
	parentID = 200004

	//linkType is the link type for the nodelink
	linkType = "FILE"

	//fileRemarks is the remarks for the file
	fileRemarks = "This is test Version"
)

//mock_insert_file mock a file insert operation in mcre
func mockInsertFile() error {

	connString := mysql.ConstructConnString(dbUserName, dbPassword, dbName)

	seqDatabase := mysql.ConnectDatabase(connString)

	//seq Repository
	seq := mysql.NewSeq(seqDatabase)

	//Generate ID for New File
	fileID, err := seq.Find("FILENAME")
	if err != nil {
		return errors.New("Error on IDGenerator for Filename " + err.Error())
	}

	fmt.Println("New FILE ID: ", fileID)

	var timeStart = time.Now().Format(dateOnly)
	var timeAfter = time.Now().Add(150 * time.Hour).Format(dateOnly)
	var fgName = strconv.Itoa(fileID) + fileExt
	var destPath = storageFolder + fgName

	nodeID, err := seq.Find("NODE") //Generate a new ID for Node
	if err != nil {
		return errors.New("Error on IDGenerator for Node " + err.Error())
	}

	fmt.Println("New Node ID: ", nodeID)

	//create a test file
	_, err = CreateFile(fileName, fileSize)
	if err != nil {
		return errors.New("Error on Create File " + err.Error())
	}

	//populate the struct with data
	newNode, err := CreateNewNode(nodeID, fileName, timeStart, graphID, fileType, userID)
	if err != nil {
		return errors.New("Error on populating Node struct " + err.Error())
	}
	newFmedia, err := CreateNewFm(nodeID, fileName, fileExt, storageFolder, fgName, fileName, fileRemarks, fileSize, 1, 1)
	if err != nil {
		return errors.New("Error on populating Fmedia struct" + err.Error())
	}
	newNL, err := CreateNewNL(parentID, nodeID, linkType)
	if err != nil {
		return errors.New("Error on populating Nodelink struct" + err.Error())
	}
	newFv, err := CreateNewFv(nodeID, timeAfter, fileRemarks, timeStart, "1", verState)
	if err != nil {
		return errors.New("Error on populating Fversion struct" + err.Error())
	}
	newConv, err := CreateNewConv(nodeID, fileExt, destPath)
	if err != nil {
		return errors.New("Error on populating Convqueue struct" + err.Error())
	}
	//Copy the file to destination
	err = CopyFile(fileName, destPath)
	if err != nil {
		return errors.New("Error on Copy File " + err.Error())
	}

	//Prepare statement for insert
	stmt := new(PreparedStmt)
	err = stmt.PrepareStatement(seqDatabase)
	if err != nil {
		return errors.New("Error on Prepare Statment " + err.Error())
	}

	//Insert data into database
	err = CommitIntoDatabase(seqDatabase, *stmt, newNode, newFmedia, newNL, newFv, newConv)
	if err != nil {
		return errors.New("Error on CommitIntoDatabase " + err.Error())
	}

	//now update the seq table's value after the file has been successfully inserted
	err = seq.Update("FILE", fileID)
	if err != nil {
		return errors.New("Error on Update File Value " + err.Error())
	}
	err = seq.Update("NODE", nodeID)
	if err != nil {
		return errors.New("Error on Update Node Value " + err.Error())
	}

	seqDatabase.Close()

	return nil
}

//CreateFile create a new file and truncate it to the file size passed in by user
func CreateFile(filename string, fileSize int64) (*os.File, error) {
	if filename == "" {
		return nil, errors.New("No Filename being supplied into the function")
	}
	newFile, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	err1 := os.Truncate(newFile.Name(), fileSize)
	if err1 != nil {
		return nil, err
	}
	newFile.Close()
	return newFile, err
}

//CreateNewNode return an instantiated Node
func CreateNewNode(nodeID int, fileName string, timeStart string, graphID int, fileType string, userID int) (Node, error) {
	if nodeID == 0 || fileName == "" || timeStart == "" || graphID == 0 || fileType == "" || userID == 0 {
		return Node{}, errors.New("Not enough argument supplied")
	}
	return Node{
		NodeID:    nodeID,
		NodeBits:  0,
		NodeDesc:  fileName,
		NodeDT:    timeStart,
		NodeGID:   graphID,
		NodeHash:  "",
		NodeLevel: -32768,
		NodeType:  fileType,
		NodeUID:   userID,
	}, nil
}

//CreateNewFm returns an instantiated Fmedia
func CreateNewFm(nodeID int, fileName string, fileExt string, storageFolder string, fgName string, foName string, fremark string, fileSize int, fileStatus int, fileType int) (Fmedia, error) {
	if nodeID == 0 || fileName == "" || fileExt == "" || storageFolder == "" || fgName == "" || foName == "" || fileSize == 0 {
		return Fmedia{}, errors.New("Not enough argument supplied")
	}

	return Fmedia{
		NodeID:   nodeID,
		FDesc:    fileName,
		FExt:     fileExt,
		FFulPath: storageFolder,
		FGName:   fgName,
		FOName:   fileName,
		FRemark:  fremark,
		FSize:    fileSize,
		FStatus:  fileStatus,
		FType:    fileType,
	}, nil
}

//CreateNewNL returns an instantiated NodeLink
func CreateNewNL(parentNodeID int, childNodeID int, linkType string) (NodeLink, error) {
	if parentNodeID == 0 || childNodeID == 0 || linkType == "" {
		return NodeLink{}, errors.New("Not enough arguments supplied")
	}
	return NodeLink{
		LinkPNodeID: parentNodeID,
		LinkCNodeID: childNodeID,
		LinkType:    linkType,
	}, nil
}

//CreateNewFv returns an instantiated Fversion
func CreateNewFv(nodeID int, timeAfter string, remarks string, timeStart string, version string, verState int) (Fverinfo, error) {
	if nodeID == 0 || timeAfter == "" || remarks == "" || timeStart == "" || version == "" || verState == 0 {
		return Fverinfo{}, errors.New("Not enough argument supplied")
	}
	return Fverinfo{
		NodeID:    nodeID,
		EndDate:   timeAfter,
		Remarks:   remarks,
		StartDate: timeStart,
		Version:   version,
		VerState:  verState,
	}, nil
}

//CreateNewConv returns an instantiated Convqueue
func CreateNewConv(nodeID int, fileExt string, destPath string) (Convqueue, error) {
	if nodeID == 0 || fileExt == "" || destPath == "" {
		return Convqueue{}, errors.New("Not enough argument supplied")
	}
	return Convqueue{
		NodeID:   nodeID,
		ConvType: "PDF",
		FExt:     fileExt,
		FFulPath: destPath,
		INSDate:  time.Now().Format(dateAndTime),
		Priority: 1,
	}, nil
}

//CopyFile will copy the file to the destination path
func CopyFile(fileName string, destPath string) error {
	oriFile, err := os.Open(fileName)
	if err != nil {
		return err
	}

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}

	bytesWritten, err := io.Copy(destFile, oriFile)
	if err != nil {
		return err
	}
	fmt.Println("Bytes copied: ", bytesWritten)

	oriFile.Close()
	destFile.Close()

	return nil
}

//PrepareNodeStmt makes a prepared statement for Node insert
func PrepareNodeStmt(db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.Prepare("INSERT node SET nodeid=?,nodebits=?,nodedesc=?,nodedt=?,nodegid=?,nodehash=?,nodelevel=?,nodetype=?,nodeuid=?")
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

//PrepareFmStmt makes a prepared statement for Fmedia insert
func PrepareFmStmt(db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.Prepare("INSERT fmedia SET nodeid=?,fdesc=?,fext=?,ffulpath=?,fgname=?,foname=?,fremark=?,fsize=?,fstatus=?,ftype=?")
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

//PrepareNLStmt makes a prepared statement for NodeLink insert
func PrepareNLStmt(db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.Prepare("INSERT nodelink SET linkcnodeid=?,linkpnodeid=?,linktype=?")
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

//PrepareFvStmt makes a prepared statement for Fversion insert
func PrepareFvStmt(db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.Prepare("INSERT fverinfo SET nodeid=?,enddate=?,remarks=?,startdate=?,version=?,verstate=?")
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

//PrepareConvStmt makes a prepared statement for Convqueue insert
func PrepareConvStmt(db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.Prepare("INSERT convqueue SET nodeid=?,convtype=?,fext=?,ffulpath=?,insdate=?,priority=?")
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

//PrepareStatement will call all Preparestatment function
func (stmt *PreparedStmt) PrepareStatement(db *sql.DB) error {
	var err error
	stmt.NodeStmt, err = PrepareNodeStmt(db)
	if err != nil {
		return err
	}
	stmt.FmediaStmt, err = PrepareFmStmt(db)
	if err != nil {
		return err
	}
	stmt.NLStmt, err = PrepareNLStmt(db)
	if err != nil {
		return err
	}
	stmt.FvStmt, err = PrepareFvStmt(db)
	if err != nil {
		return err
	}
	stmt.ConvStmt, err = PrepareConvStmt(db)
	if err != nil {
		return err
	}
	return nil
}

//CommitIntoDatabase will insert Node,Fmedia,Nodelink,Fversion and convqueue into database
func CommitIntoDatabase(db *sql.DB, stmt PreparedStmt, newNode Node, newFmedia Fmedia, newNL NodeLink, newFv Fverinfo, newConv Convqueue) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Stmt(stmt.NodeStmt).Exec(newNode.NodeID, newNode.NodeBits, newNode.NodeDesc, newNode.NodeDT, newNode.NodeGID, newNode.NodeHash, newNode.NodeLevel, newNode.NodeType, newNode.NodeUID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Stmt(stmt.FmediaStmt).Exec(newFmedia.NodeID, newFmedia.FDesc, newFmedia.FExt, newFmedia.FFulPath, newFmedia.FGName, newFmedia.FOName, newFmedia.FRemark, newFmedia.FSize, newFmedia.FStatus, newFmedia.FType)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Stmt(stmt.NLStmt).Exec(newNL.LinkCNodeID, newNL.LinkPNodeID, newNL.LinkType)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Stmt(stmt.FvStmt).Exec(newFv.NodeID, newFv.EndDate, newFv.Remarks, newFv.StartDate, newFv.Version, newFv.VerState)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Stmt(stmt.ConvStmt).Exec(newConv.NodeID, newConv.ConvType, newConv.FExt, newConv.FFulPath, newConv.INSDate, newConv.Priority)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
