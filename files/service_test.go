package files

import (
	"os"
	"testing"

	"github.com/go-panton/mcre/infra/store/mysql"
)

func TestDownload(t *testing.T) {
}

func TestCreateFile(t *testing.T) {
	//Create a temp file in test directory
	tempDir := os.TempDir()
	newFile, err := os.Create(tempDir + "\\test.txt")
	if err != nil {
		t.Errorf("Error Creating a new file: %v", err.Error())
		t.FailNow()
	}
	err = newFile.Truncate(200)
	if err != nil {
		t.Errorf("Error truncating the file with dummy data: %v", err.Error())
	}

	seqs := mysql.NewMockSeqRepository()
	nodes := mysql.NewMockNodeRepository()
	fmedias := mysql.NewMockFmediaRepository()
	fverinfos := mysql.NewMockFverinfoRepository()
	nls := mysql.NewMockNodelinkRepository()
	convs := mysql.NewMockConvqueueRepository()
	fmps := mysql.NewMockFmpolicyRepository()

	svc := NewService(seqs, nodes, fmedias, nls, fverinfos, convs, fmps)

	filename := newFile.Name()

	// close the file because the service does not need the file pointer anymore
	err = newFile.Close()
	if err != nil {
		t.Errorf("Error closing file: %v", err.Error())
		t.FailNow() //cant close file means cant delete so just fail the test straight
	}

	err = svc.Create(filename)
	if err != nil {
		t.Errorf("Error while calling create service: %v", err.Error())
	}

	err = os.Remove(tempDir + "\\test.txt")
	if err != nil {
		t.Errorf("Error removing test file: %v", err.Error())
		t.FailNow() // can't remove just fail the test straight
	}
}

func TestCreateVersion(t *testing.T) {
	//Create a temp file in test directory
	tempDir := os.TempDir()
	newFile, err := os.Create(tempDir + "\\test.txt")
	if err != nil {
		t.Errorf("Error Creating a new file: %v", err.Error())
		t.FailNow()
	}
	err = newFile.Truncate(200)
	if err != nil {
		t.Errorf("Error truncating the file with dummy data: %v", err.Error())
	}

	seqs := mysql.NewMockSeqRepository()
	nodes := mysql.NewMockNodeRepository()
	fmedias := mysql.NewMockFmediaRepository()
	fverinfos := mysql.NewMockFverinfoRepository()
	nls := mysql.NewMockNodelinkRepository()
	convs := mysql.NewMockConvqueueRepository()
	fmps := mysql.NewMockFmpolicyRepository()

	svc := NewService(seqs, nodes, fmedias, nls, fverinfos, convs, fmps)

	filename := newFile.Name()

	currFNodeID := 200264

	// close the file because the service does not need the file pointer anymore
	err = newFile.Close()
	if err != nil {
		t.Errorf("Error closing file: %v", err.Error())
		t.FailNow() //cant close file means cant delete so just fail the test straight
	}
	err = svc.Version(filename, currFNodeID)
	if err != nil {
		t.Errorf("Error while running version service: %v", err.Error())
	}

	err = os.Remove(tempDir + "\\test.txt")
	if err != nil {
		t.Errorf("Error removing test file: %v", err.Error())
		t.FailNow() // can't remove just fail the test straight
	}
}
