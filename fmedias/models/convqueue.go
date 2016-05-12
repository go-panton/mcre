package models

import "errors"

//Convqueue struct defie the data of convqueue that gets store in the database
type Convqueue struct {
	NodeID   int
	Convtype string
	FExt     string
	FFulpath string
	InsDate  string //t.format("2006-01-02 15:04:05")
	Priority int
}

//ConvqueueRepository interface define the interface methods that can be accessed by others
type ConvqueueRepository interface {
	Insert(Convqueue) error
	Delete(int) error
	GetInsertStr(Convqueue) (string, error)
	GetDeleteStr(int) (string, error)
}

//NewConvqueue return a new COnvqueue struct based on parameters provided
func NewConvqueue(nodeID int, insDate string, fExt string, fFulpath string) (Convqueue, error) {
	if nodeID == 0 || insDate == "" || fExt == "" || fFulpath == "" {
		return Convqueue{}, errors.New("New convqueue parameter cannot be empty")
	}
	return Convqueue{
		NodeID:   nodeID,
		Convtype: "PDF", //PDF is the default type to be used
		FExt:     fExt,
		FFulpath: fFulpath,
		InsDate:  insDate,
		Priority: 1,
	}, nil
}
