package fmedias

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-panton/mcre/fmedias/models"
	"github.com/go-panton/mcre/infra/store/mysql"
)

func TestFmpInsert(t *testing.T) {
	fmps := mysql.NewMockFmpolicyRepository()

	tests := []struct {
		Fmpolicy models.Fmpolicy
		Want     error
	}{
		{models.Fmpolicy{FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1, NodeID: 212121}, nil},
		{models.Fmpolicy{FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1}, errors.New("NodeID cannot be empty or 0")},
	}
	for _, test := range tests {
		err := fmps.Insert(test.Fmpolicy)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestFmpUpdate(t *testing.T) {
	fmps := mysql.NewMockFmpolicyRepository()

	tests := []struct {
		Fmpolicy models.Fmpolicy
		Want     error
	}{
		{models.Fmpolicy{FmpID: 121, FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1, NodeID: 212121}, nil},
		{models.Fmpolicy{FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1, NodeID: 212121}, errors.New("Parameter cannot be empty or 0")},
		{models.Fmpolicy{FmpID: 121, FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1}, errors.New("Parameter cannot be empty or 0")},
	}

	for _, test := range tests {
		err := fmps.Update(test.Fmpolicy)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestFmpDelete(t *testing.T) {
	fmps := mysql.NewMockFmpolicyRepository()

	tests := []struct {
		fmpID int
		Want  error
	}{
		{123, nil},
		{0, errors.New("Fmp ID cannot be empty or 0")},
	}
	for _, test := range tests {
		err := fmps.Delete(test.fmpID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestFmpFind(t *testing.T) {
	fmps := mysql.NewMockFmpolicyRepository()

	tests := []struct {
		fmpID int
		Want  error
	}{
		{121, nil},
		{0, errors.New("Fmp ID cannot be empty or 0")},
	}

	for _, test := range tests {
		fmpolicy, err := fmps.Find(test.fmpID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Fmpolicy Find: ", fmpolicy)
		}
	}
}

func TestFmpFindUsingNodeID(t *testing.T) {
	fmps := mysql.NewMockFmpolicyRepository()

	tests := []struct {
		NodeID int
		Want   error
	}{
		{232323, nil},
		{0, errors.New("NodeID cannot be empty or 0")},
	}

	for _, test := range tests {
		res, err := fmps.FindUsingNodeID(test.NodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Fmp FindN: ", res)
		}
	}
}

func TestFmpGetInsertStr(t *testing.T) {
	fmps := mysql.NewMockFmpolicyRepository()

	tests := []struct {
		Fmpolicy models.Fmpolicy
		Want     error
	}{
		{models.Fmpolicy{FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1, NodeID: 212121}, nil},
		{models.Fmpolicy{FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1}, errors.New("Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		insertStr, err := fmps.GetInsertStr(test.Fmpolicy)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Fmp Insert Str: ", insertStr)
		}
	}
}

func TestFmpGetUpdateStr(t *testing.T) {
	fmps := mysql.NewMockFmpolicyRepository()

	tests := []struct {
		Fmpolicy models.Fmpolicy
		Want     error
	}{
		{models.Fmpolicy{FmpID: 122, FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1, NodeID: 212121}, nil},
		{models.Fmpolicy{FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1, NodeID: 212121}, errors.New("Parameter cannot be empty or 0")},
		{models.Fmpolicy{FmpID: 122, FmpDownload: 1, FmpRevise: 1, FmpView: 1, FmpUGID: 223232, FmpUGType: 1}, errors.New("Parameter cannot be empty or 0")},
	}
	for _, test := range tests {
		updateStr, err := fmps.GetUpdateStr(test.Fmpolicy)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Fmp Update Str: ", updateStr)
		}
	}
}

func TestFmpGetDeleteStr(t *testing.T) {
	fmps := mysql.NewMockFmpolicyRepository()

	tests := []struct {
		fmpID int
		Want  error
	}{
		{123, nil},
		{0, errors.New("Fmp ID cannot be empty or 0")},
	}
	for _, test := range tests {
		deleteStr, err := fmps.GetDeleteStr(test.fmpID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Fmp Delete Str: ", deleteStr)
		}
	}
}
