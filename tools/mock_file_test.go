package tools

import "testing"

type FileTest struct {
	FileName string
	FileSize int64
}

type NodeTest struct {
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

var testCaseFile = []FileTest{
	{"Rawr.pdf", 500},
	{"I.gif", 300},
	{"am.jpeg", 0},
	{"a.zxc", 500},
	{FileName: "tiger"},
	{FileSize: 200},
	{},
}

var testCaseNode = []NodeTest{
	{12312},
}

func TestCreateFile(t *testing.T) {
	for _, test := range testCaseFile {
		newfile, err := CreateFile(test.FileName, test.FileSize)
		if err != nil {
			if newfile != nil {
				t.Errorf("Expect no file being create, but newfile is not nil")
			}
		}
	}
}

func TestCreateNewNode(t *testing.T) {

}

//func TestMockFile(t *testing.T) {
//	err := mock_insert_file()
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//}
