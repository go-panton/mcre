package mysql

import (
	"os"
	"strings"
	"testing"
	"time"

	files "github.com/go-panton/mcre/files"
	fmedias "github.com/go-panton/mcre/fmedias/models"
	nodes "github.com/go-panton/mcre/nodes/models"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const fakeDBName = "foo"

// func TestUserRepo(t *testing.T) {
// 	db, err := sql.Open("mysql", "root:root123@/go_panton")
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
//
// 	NewUser(db).Insert("test", "random")
//
// 	result, err := NewUser(db).Find("test")
// 	if err != nil || result == nil {
// 		t.Errorf("No result from database")
// 	}
// 	fmt.Println(result)
// }
//
// func TestSeqFind(t *testing.T) {
// 	tests := []struct {
// 		Key   string
// 		Want  int
// 		Error error
// 	}{
// 		{"NODE", 200905, nil},
// 		{"FILENAME", 766, nil},
// 		{"", 0, errors.New("The key is empty.")},
// 	}
//
// 	db, err := sql.Open("mysql", "root:root123@/ptd_new")
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	for _, test := range tests {
// 		res, err := NewSeq(db).Find(test.Key)
// 		if err != nil {
// 			if err.Error() != test.Error.Error() {
// 				t.Errorf("Got: %v, Want: %v", err.Error(), test.Error.Error())
// 			}
// 		}
// 		if res != test.Want {
// 			t.Errorf("Got: %v,Want: %v", res, test.Want)
// 		}
// 	}
// }

// func TestSeqUpdate(t *testing.T) {
// 	db, err := sql.Open("mysql", "root:root123@/ptd_new")
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
//
// 	tests := []struct {
// 		Key   string
// 		Value int
// 		Error error
// 	}{
// 		{"NODE", 200905, nil},
// 		{"FILENAME", 766, nil},
// 		{"NODE", 0, errors.New("Update Value cannot be less than 1")},
// 		{"FILENAME", 0, errors.New("Update Value cannot be less than 1")},
// 		{"", 0, errors.New("The key is empty.")},
// 	}
//
// 	for _, test := range tests {
// 		err := NewSeq(db).Update(test.Key, test.Value)
//
// 		if err != nil {
// 			if err.Error() != test.Error.Error() {
// 				t.Errorf("Got: %v, Want: %v", err.Error(), test.Error.Error())
// 			}
// 		}
// 	}
// }

// func TestNewDB(t *testing.T) {
// 	db, err := sql.Open("mysql", "root:root123@/")
// 	if err != nil {
// 		t.Fatalf("Open: %v", err)
// 	}
//
// 	defer closeDB(t, db)
//
// 	createDB(t, db)
//
// 	useDB(t, db)
//
// 	createTable(t, db)
//
// }

//TestInsert test on the isnert using dummy data *Not using transaction
// func TestInsert(t *testing.T) {
// 	db, err := sqlx.Open("mysql", "root:root123@/")
// 	if err != nil {
// 		t.Fatalf("Open: %v", err)
// 	}
//
// 	defer closeDB(t, db)
//
// 	useDB(t, db)
//
//  insertSeq(t, db)
//
//  insertNode(t, db)
//
//	insertFmedia(t, db)
//
// 	insertFver(t, db)
//
// 	insertFmp(t, db)
//
// 	insertNl(t, db)
// }

// func TestCreateService(t *testing.T) {
// 	db, err := sqlx.Open("mysql", "root:root123@/")
// 	if err != nil {
// 		t.Fatalf("Open: %v", err)
// 	}
//
// 	useDB(t, db)
//
// 	defer closeDB(t, db)
//
// 	tempDir := os.TempDir()
// 	newFile, err := os.Create(tempDir + "\\test.txt")
// 	if err != nil {
// 		t.Errorf("Error Creating a new file: %v", err.Error())
// 		t.FailNow()
// 	}
//
// 	defer newFile.Close()
//
// 	err = newFile.Truncate(200)
// 	if err != nil {
// 		t.Errorf("Error truncating the file with dummy data: %v", err.Error())
// 	}
//
// 	fmediaR := NewFmedia(db)
// 	fverR := NewFverinfo(db)
// 	fmpR := NewFmpolicy(db)
// 	nodeR := NewNode(db)
// 	nlR := NewNodeLink(db)
// 	seqR := NewSeq(db)
// 	convR := NewConvqueue(db)
//
// 	svc := files.NewService(seqR, nodeR, fmediaR, nlR, fverR, convR, fmpR)
//
// 	err = svc.Create(newFile.Name())
// 	if err != nil {
// 		t.Fatalf("svc.Create: %v", err)
// 	}
// }

func TestVersionService(t *testing.T) {
	db, err := sqlx.Open("mysql", "root:root123@/")
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	useDB(t, db)

	defer closeDB(t, db)

	db.MapperFunc(strings.ToUpper)

	tempDir := os.TempDir()
	newFile, err := os.Create(tempDir + "\\test.txt")
	if err != nil {
		t.Errorf("Error Creating a new file: %v", err.Error())
		t.FailNow()
	}

	defer newFile.Close()

	err = newFile.Truncate(200)
	if err != nil {
		t.Errorf("Error truncating the file with dummy data: %v", err.Error())
	}

	fmediaR := NewFmedia(db)
	fverR := NewFverinfo(db)
	fmpR := NewFmpolicy(db)
	nodeR := NewNode(db)
	nlR := NewNodeLink(db)
	seqR := NewSeq(db)
	convR := NewConvqueue(db)

	svc := files.NewService(seqR, nodeR, fmediaR, nlR, fverR, convR, fmpR)

	if err = svc.Version(newFile.Name(), 200068); err != nil {
		t.Fatalf("svc Version: %v", err)
	}
}

//exec execute the query with the arguments and will fail the test immediately if the query throwback an error
func exec(t testing.TB, db *sqlx.DB, query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		t.Fatalf("Exec of %q: %v", query, err)
	}
}

// createDB create a new DB using the fakeDBName if it does not exist
func createDB(t testing.TB, db *sqlx.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS" + fakeDBName)
	if err != nil {
		t.Fatalf("Error creating database: %v", err)
	}
}

//useDB make the db conenction to use that database
func useDB(t testing.TB, db *sqlx.DB) {
	_, err := db.Exec("USE " + fakeDBName)
	if err != nil {
		t.Fatalf("Error using database: %v", err)
	}
}

// createTable create the table for node, fmedia, fverinfo, seqtbl, fmpolicy, nodelink, convqueue if it does not exist in the database
func createTable(t testing.TB, db *sqlx.DB) {
	exec(t, db, "CREATE TABLE IF NOT EXISTS `node`(`NODEID` bigint(20) NOT NULL PRIMARY KEY,"+
		"`NDOEBITS` decimal(19,2),`NODEDESC` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin,`NODEDT` date NOT NULL,`NODEGID` bigint(20),"+
		"`NODEHASH` varchar(255),`NODELEVEL` smallint(6),`NODETYPE` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin,`NODEUID` bigint(20))")

	exec(t, db, "CREATE TABLE IF NOT EXISTS `fmedia`(`NODEID` bigint(20) NOT NULL PRIMARY KEY,"+
		"`FDESC` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin,`FEXT` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,"+
		"`FFULPATH` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,"+
		"`FGNAME` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,`FONAME` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,"+
		"`FREMARK` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin,`FSIZE` bigint(20) NOT NULL,`FSTATUS` int(11),`FTYPE` int(11))")

	exec(t, db, "CREATE TABLE IF NOT EXISTS `fverinfo`(`NODEID` bigint(20) NOT NULL PRIMARY KEY,"+
		"`ENDDATE` date,`REMARKS` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin,`STARTDATE` date NOT NULL,"+
		"`VERSION` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL, `VERSTATE` int(11) NOT NULL)")

	exec(t, db, "CREATE TABLE IF NOT EXISTS `seqtbl`(`PNAME` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,"+
		"`PSEQ` bigint(20) NOT NULL)")

	exec(t, db, "CREATE TABLE IF NOT EXISTS `fmpolicy`(`FMPID` bigint(20) NOT NULL PRIMARY KEY AUTO_INCREMENT,"+
		"`FMPDOWNLOAD` tinyint(1) NOT NULL,`FMPREVISE` tinyint(1) NOT NULL,`FMPUGID` bigint(20) NOT NULL,"+
		"`FMPUGTYPE` int(11) NOT NULL,`FMPVIEW` tinyint(1) NOT NULL,`NODEID` bigint(20) NOT NULL, FOREIGN KEY(NODEID) REFERENCES node(NODEID))")

	exec(t, db, "CREATE TABLE IF NOT EXISTS `nodelink`(`LINKCNODEID` bigint(20) NOT NULL PRIMARY KEY,"+
		"`LINKPNODEID` bigint(20) NOT NULL,`LINKTYPE` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin)")

	exec(t, db, "CREATE TABLE IF NOT EXISTS `convqueue`(`NODEID` bigint(20) NOT NULL PRIMARY KEY,"+
		"`CONVTYPE` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,`FEXT` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,"+
		"`FFULPATH` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL, INSDATE datetime NOT NULL, PRIORITY int(11) NOT NULL)")
}

// insertNode insert dummy data into the node table
func insertNode(t testing.TB, db *sqlx.DB) {
	nodes := []nodes.Node{
		nodes.Node{NodeID: 200056, NodeDesc: "test", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1},
		nodes.Node{NodeID: 200057, NodeDesc: "test1", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1},
		nodes.Node{NodeID: 200058, NodeDesc: "test2", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1},
	}

	for _, node := range nodes {
		exec(t, db, "INSERT node SET nodeid=?, nodedesc=?, nodedt=?, nodegid=?, nodetype=?, nodeuid=?", node.NodeID, node.NodeDesc, node.NodeDT, node.NodeGID, node.NodeType, node.NodeUID)
	}
}

// insertFmedia insert dummy data into fmedia table
func insertFmedia(t testing.TB, db *sqlx.DB) {
	fmediaStruct := []fmedias.Fmedia{
		fmedias.Fmedia{NodeID: 200056, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test.txt", FRemark: "Testing", FSize: 43432, FStatus: 1, FType: 1},
		fmedias.Fmedia{NodeID: 200057, FDesc: "test1.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "346.txt", FOName: "test1.txt", FRemark: "Testing 1", FSize: 43432, FStatus: 1, FType: 1},
		fmedias.Fmedia{NodeID: 200058, FDesc: "test2.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "347.txt", FOName: "test2.txt", FRemark: "Testing 2", FSize: 43432, FStatus: 1, FType: 1},
	}

	for _, fmedia := range fmediaStruct {
		exec(t, db, "INSERT fmedia SET nodeid=?, fdesc=?, fext=?,ffulpath=?, fgname=?, foname=?, fremark=?, fsize=?, fstatus=?, ftype=?", fmedia.NodeID, fmedia.FDesc, fmedia.FExt,
			fmedia.FFulPath, fmedia.FGName, fmedia.FOName, fmedia.FRemark, fmedia.FSize, fmedia.FStatus, fmedia.FType)
	}
}

// insertFver insert dummy data into fverinfo table
func insertFver(t testing.TB, db *sqlx.DB) {
	fvers := []fmedias.Fverinfo{
		fmedias.Fverinfo{NodeID: 200056, EndDate: "2016-10-10", Remarks: "This is auto-generated version.", StartDate: "2016-05-01", Version: "1", VerState: 1},
		fmedias.Fverinfo{NodeID: 200057, EndDate: "2016-10-10", Remarks: "This is auto-generated version.", StartDate: "2016-05-01", Version: "1", VerState: 1},
		fmedias.Fverinfo{NodeID: 200058, EndDate: "2016-10-10", Remarks: "This is auto-generated version.", StartDate: "2016-05-01", Version: "1", VerState: 1},
	}

	for _, fver := range fvers {
		exec(t, db, "INSERT Fverinfo SET nodeid=?, enddate=?, remarks=?, startdate=?, version=?, verstate=?", fver.NodeID, fver.EndDate, fver.Remarks, fver.StartDate,
			fver.Version, fver.VerState)
	}
}

// insertFmp insert dummy data into fmpolicy table
func insertFmp(t testing.TB, db *sqlx.DB) {
	fmps := []fmedias.Fmpolicy{
		fmedias.Fmpolicy{FmpID: 1, FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 0, FmpUGType: 0, NodeID: 200056},
		fmedias.Fmpolicy{FmpID: 2, FmpDownload: 1, FmpRevise: 0, FmpView: 1, FmpUGID: 1, FmpUGType: 1, NodeID: 200057},
		fmedias.Fmpolicy{FmpID: 3, FmpDownload: 0, FmpRevise: 1, FmpView: 1, FmpUGID: 1, FmpUGType: 1, NodeID: 200058},
	}

	for _, fmp := range fmps {
		exec(t, db, "INSERT fmpolicy SET fmpid=?, fmpdownload=?, fmprevise=?, fmpugid=?, fmpugtype=?, fmpview=?, nodeid=?", fmp.FmpID, fmp.FmpDownload, fmp.FmpRevise, fmp.FmpUGID,
			fmp.FmpUGType, fmp.FmpView, fmp.NodeID)
	}
}

// insertNl insert dummy data into nodelink table
func insertNl(t testing.TB, db *sqlx.DB) {
	nls := []nodes.Nodelink{
		nodes.Nodelink{LinkCNodeID: 200056, LinkPNodeID: 200004, LinkType: "FILE"},
		nodes.Nodelink{LinkCNodeID: 200057, LinkPNodeID: 200004, LinkType: "FILE"},
		nodes.Nodelink{LinkCNodeID: 200058, LinkPNodeID: 200004, LinkType: "FILE"},
	}

	for _, nl := range nls {
		exec(t, db, "INSERT nodelink SET linkcnodeid=?, linkpnodeid=?, linktype=?", nl.LinkCNodeID, nl.LinkPNodeID, nl.LinkType)
	}
}

// insertSeq insert NODE and FILENAME with its value into seqtbl table
func insertSeq(t testing.TB, db *sqlx.DB) {
	exec(t, db, "INSERT seqtbl SET PNAME=?, PSEQ=?", "NODE", 200058)
	exec(t, db, "INSERT seqtbl SET PNAME=?, PSEQ=?", "FILENAME", 347)
}

// drop database drop the whole database including it's data and table definition
func dropDatabase(t testing.TB, db *sqlx.DB) {
	if _, err := db.Exec("DROP DATABASE " + fakeDBName); err != nil {
		t.Fatalf("Exec Wipe: %v", err)
	}
}

// closeDB close the db connection
func closeDB(t testing.TB, db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		t.Fatalf("Error Closing DB: %v", err)
	}
}
