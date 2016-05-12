package models

import "errors"

//Fmpolicy struct defines the data of fmpolicy that stores in database
//
//Do note that only FmpID and NodeID can never be 0, but fmpid is auto incremented in database so insert
//does not require fmpid to be set beforehand
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
	// Insert insert the fmpolicy into database based on the fmpolicy provided
	//
	// Returns error when :-
	// - insert failed
	// - Failed preparing statement
	Insert(Fmpolicy) error
	// Update update the value of fmpolicy in the database based on the fmpolicy provided
	//
	// Return error when :-
	// - update failed
	// - failed preparing statement
	Update(Fmpolicy) error
	// Delete delete the fmpolicy from database based on the fmp ID provided
	//
	// Return error when :-
	// - delete failed
	// - failed preparing statement
	Delete(int) error
	// Find return a fmpolicy object based on the fmp ID provided
	//
	// Return error when :-
	// - fmp ID is empty or 0
	// - No record correspond to the fmp ID provided
	// - Invalid query
	// - find failed
	Find(int) (Fmpolicy, error)
	// FindUsingNodeID return a slice of fmpolicy object based on the node ID provided
	//
	// Return error when :-
	// - node ID is empty or 0
	// - No record correspond to the fmp ID provided
	// - Invalid query
	// - find failed
	FindUsingNodeID(int) ([]Fmpolicy, error)
	// GetInsertStr return a string of insert statement based on the fmpolicy provided
	//
	// Return error when :-
	// - node ID of the fmpolicy provided is empty or 0
	GetInsertStr(Fmpolicy) (string, error)
	// GetUpdateStr return a string of update statement based on the fmpolicy provided
	//
	// Return error when :-
	// - node ID of the fmpolicy provided is empty or 0
	GetUpdateStr(Fmpolicy) (string, error)
	// GetDeleteStr return a string of delete statement based on the fmp ID provided
	//
	// Return error when :-
	// - fmp ID provided is empty or 0
	GetDeleteStr(int) (string, error)
}

// NewFmpolicy returns a Fmpolicy based on the parameters passed in
//
// Returns error when :-
// - NodeID is empty or 0
func NewFmpolicy(fmpdownload, fmprevise, fmpview, fmpugid, fmpugtype, nodeID int) (Fmpolicy, error) {
	if nodeID == 0 {
		return Fmpolicy{}, errors.New("New Fmpolicy Node ID cannot be empty")
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
