package models

import "errors"

//Fmpolicy struct defines the data of fmpolicy that stores in database
type Fmpolicy struct {
	FmpID       int
	FmpDownload int
	FmpRevise   int
	FmpUGID     int
	FmpUGType   int
	FmpView     int
	NodeID      int
}

//FmpolicyRepository interface defines the interface methods
type FmpolicyRepository interface {
	Insert(Fmpolicy) error
	Update(Fmpolicy) error
	Delete(int) error
	Find(int) (Fmpolicy, error)
	GetInsertStr(Fmpolicy) (string, error)
	GetUpdateStr(Fmpolicy) (string, error)
	GetDeleteStr(int) (string, error)
}

func NewFmpolicy(fmpdownload, fmprevise, fmpview, fmpugid, fmpugtype, nodeID int) (Fmpolicy, error) {
	if nodeID == 0 {
		return Fmpolicy{}, errors.New("Node ID cannot be empty")
	}
	return Fmpolicy{
		FmpDownload: fmpdownload,
		FmpRevise:   fmprevise,
		FmpView:     fmpview,
		FmpUGID:     fmpugid,
		FmpUGType:   fmpugtype,
		NodeID:      nodeID,
	}, nil
}
