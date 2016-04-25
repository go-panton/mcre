package fmedias

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-panton/mcre/fmedias/models"
	"github.com/go-panton/mcre/infra/store/mysql"
)

func TestInsert(t *testing.T) {
	fmedias := mysql.NewMockFmediaRepository()

	tests := []struct {
		Fmedia models.Fmedia
		Want   error
	}{
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, nil},
		{models.Fmedia{FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
	}
	for _, test := range tests {
		err := fmedias.Insert(test.Fmedia)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestUpdate(t *testing.T) {
	fmedias := mysql.NewMockFmediaRepository()

	tests := []struct {
		Fmedia models.Fmedia
		Want   error
	}{
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, nil},
		{models.Fmedia{FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
	}
	for _, test := range tests {
		err := fmedias.Update(test.Fmedia)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestDelete(t *testing.T) {
	fmedias := mysql.NewMockFmediaRepository()

	tests := []struct {
		NodeID int
		Want   error
	}{
		{24242, nil},
		{0, errors.New("Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		err := fmedias.Delete(test.NodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestFind(t *testing.T) {
	fmedias := mysql.NewMockFmediaRepository()

	tests := []struct {
		NodeID int
		Want   error
	}{
		{24242, nil},
		{0, errors.New("Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		fmedia, err := fmedias.Find(test.NodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result of find: ", fmedia)
		}
	}
}

func TestFindByFileDesc(t *testing.T) {
	fmedias := mysql.NewMockFmediaRepository()

	tests := []struct {
		FileDesc string
		Want     error
	}{
		{"test.txt", nil},
		{"", errors.New("File description cannot be empty")},
	}

	for _, test := range tests {
		res, err := fmedias.FindByFileDesc(test.FileDesc)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			for _, fmedia := range res {
				fmt.Println("Result slice: ", fmedia)
			}
		}
	}
}

func TestGetInsertStr(t *testing.T) {
	fmedias := mysql.NewMockFmediaRepository()

	tests := []struct {
		Fmedia models.Fmedia
		Want   error
	}{
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, nil},
		{models.Fmedia{FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
	}
	for _, test := range tests {
		insertStr, err := fmedias.GetInsertStr(test.Fmedia)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Insert string: ", insertStr)
		}
	}
}

func TestGetUpdateStr(t *testing.T) {
	fmedias := mysql.NewMockFmediaRepository()

	tests := []struct {
		Fmedia models.Fmedia
		Want   error
	}{
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, nil},
		{models.Fmedia{FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FOName: "test1.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FRemark: "Randomly Testing", FSize: 43432, FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
		{models.Fmedia{NodeID: 23231, FDesc: "test.txt", FExt: ".txt", FFulPath: "D:\\PantonSys\\PTD\\ISOStorage\\1\\", FGName: "345.txt", FOName: "test1.txt", FRemark: "Randomly Testing", FStatus: 1, FType: 1}, errors.New("Paramater cannot be empty")},
	}

	for _, test := range tests {
		res, err := fmedias.GetUpdateStr(test.Fmedia)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v ", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Update String: ", res)
		}
	}
}

func TestGetDeleteStr(t *testing.T) {
	fmedias := mysql.NewMockFmediaRepository()

	tests := []struct {
		NodeID int
		Want   error
	}{
		{24242, nil},
		{0, errors.New("Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		deleteStr, err := fmedias.GetDeleteStr(test.NodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Delete string: ", deleteStr)
		}
	}
}
