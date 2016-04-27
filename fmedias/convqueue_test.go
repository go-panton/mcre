package fmedias

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-panton/mcre/fmedias/models"
	"github.com/go-panton/mcre/infra/store/mysql"
)

func TestConvqueueInsert(t *testing.T) {
	convqueue := mysql.NewMockConvqueueRepository()

	tests := []struct {
		Convqueue models.Convqueue
		Want      error
	}{
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FExt: ".txt", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10", Priority: 1}, nil},
		{models.Convqueue{Convtype: "PDF", FExt: ".txt", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, FExt: ".txt", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FExt: ".txt", InsDate: "2016-01-07 12:26:10", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FExt: ".txt", FFulpath: "D:/PantonSys/", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FExt: ".txt", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10"}, errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		err := convqueue.Insert(test.Convqueue)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestConvqueueDelete(t *testing.T) {
	convqueue := mysql.NewMockConvqueueRepository()

	tests := []struct {
		nodeID int
		Want   error
	}{
		{232322, nil},
		{0, errors.New("Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		err := convqueue.Delete(test.nodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v , Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestConvqueueGetInsertStr(t *testing.T) {
	convqueue := mysql.NewMockConvqueueRepository()

	tests := []struct {
		convqueue models.Convqueue
		Want      error
	}{
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FExt: ".txt", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10", Priority: 1}, nil},
		{models.Convqueue{Convtype: "PDF", FExt: ".txt", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, FExt: ".txt", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FExt: ".txt", InsDate: "2016-01-07 12:26:10", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FExt: ".txt", FFulpath: "D:/PantonSys/", Priority: 1}, errors.New("Parameter cannot be empty")},
		{models.Convqueue{NodeID: 232323, Convtype: "PDF", FExt: ".txt", FFulpath: "D:/PantonSys/", InsDate: "2016-01-07 12:26:10"}, errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		insertStr, err := convqueue.GetInsertStr(test.convqueue)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v , Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Insert String: ", insertStr)
		}
	}
}

func TestConvqueueGetDeleteStr(t *testing.T) {
	convqueue := mysql.NewMockConvqueueRepository()

	tests := []struct {
		nodeID int
		Want   error
	}{
		{232322, nil},
		{0, errors.New("Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		deleteStr, err := convqueue.GetDeleteStr(test.nodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Delete String: ", deleteStr)
		}
	}
}
