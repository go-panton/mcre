package fmedias

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-panton/mcre/fmedias/models"
	"github.com/go-panton/mcre/infra/store/mysql"
)

func TestFverinfoInsert(t *testing.T) {
	fverinfos := mysql.NewMockFverinfoRepository()

	tests := []struct {
		Fverinfo models.Fverinfo
		Want     error
	}{
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1", VerState: 1}, nil},
		{models.Fverinfo{EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, Remarks: "This is a test", Version: "v1", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1"}, errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		err := fverinfos.Insert(test.Fverinfo)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestFverinfoUpdate(t *testing.T) {
	fverinfos := mysql.NewMockFverinfoRepository()

	tests := []struct {
		Fverinfo models.Fverinfo
		Want     error
	}{
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1", VerState: 1}, nil},
		{models.Fverinfo{EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, Remarks: "This is a test", Version: "v1", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1"}, errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		err := fverinfos.Update(test.Fverinfo)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestFverinfoDelete(t *testing.T) {
	fverinfos := mysql.NewMockFverinfoRepository()

	tests := []struct {
		NodeID int
		Want   error
	}{
		{232323, nil},
		{0, errors.New("Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		err := fverinfos.Delete(test.NodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestFverinfoFind(t *testing.T) {
	fverinfos := mysql.NewMockFverinfoRepository()

	tests := []struct {
		NodeID int
		Want   error
	}{
		{232323, nil},
		{0, errors.New("Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		fverinfo, err := fverinfos.Find(test.NodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result fverinfo: ", fverinfo)
		}
	}
}

func TestFverinfoGetInsertStr(t *testing.T) {
	fverinfos := mysql.NewMockFverinfoRepository()

	tests := []struct {
		Fverinfo models.Fverinfo
		Want     error
	}{
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1", VerState: 1}, nil},
		{models.Fverinfo{EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, Remarks: "This is a test", Version: "v1", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1"}, errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		res, err := fverinfos.GetInsertStr(test.Fverinfo)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Fverinfo Insert string: ", res)
		}
	}
}

func TestFverinfoGetUpdateStr(t *testing.T) {
	fverinfos := mysql.NewMockFverinfoRepository()

	tests := []struct {
		Fverinfo models.Fverinfo
		Want     error
	}{
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1", VerState: 1}, nil},
		{models.Fverinfo{EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, Remarks: "This is a test", Version: "v1", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", VerState: 1}, errors.New("Parameter cannot be empty")},
		{models.Fverinfo{NodeID: 232323, EndDate: "2016-09-25", Remarks: "This is a test", StartDate: "2016-04-25", Version: "v1"}, errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		res, err := fverinfos.GetUpdateStr(test.Fverinfo)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Fverinfo Update string: ", res)
		}
	}
}

func TestFverinfoGetDeleteStr(t *testing.T) {
	fverinfos := mysql.NewMockFverinfoRepository()

	tests := []struct {
		NodeID int
		Want   error
	}{
		{232323, nil},
		{0, errors.New("Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		res, err := fverinfos.GetDeleteStr(test.NodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Fverinfo Delete string: ", res)
		}
	}
}
