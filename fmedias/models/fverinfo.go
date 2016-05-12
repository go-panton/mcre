package models

import "errors"

//Fverinfo struct define the data of fverinfo that stores in database
type Fverinfo struct {
	NodeID    int
	EndDate   string //t.format("2006-01-02")
	Remarks   string
	StartDate string //t.format("2006-01-02")
	Version   string
	VerState  int
}

//FverinfoRepository interface define the interface methods
type FverinfoRepository interface {
	// Insert insert the fverinfo into database based on the fverinfo struct provided
	//
	// Return error when:-
	// - Insert failed
	// - Failed preparing statement
	Insert(Fverinfo) error
	// Update update the fverinfo records in the database based on the fverinfo struct provided
	//
	// Returns error when:-
	// - Update failed
	// - Failed preparing update statement
	Update(Fverinfo) error
	// Delete delete the fverinfo records in the database based on the node ID provided
	//
	// Return error when:-
	// - Delete failed
	// - Failed preparing delete statement
	Delete(int) error
	// Find return the fverinfo records from database based on node ID provided
	//
	// Returns error when:-
	// - Find failed
	// - No such record in the database
	// - Database query error
	Find(int) (Fverinfo, error)
	// GetInsertStr return an insert statement string that can later be used as a transaction's prepared statement
	//
	// Return error when:-
	// - Parameter is empty or missing
	GetInsertStr(Fverinfo) (string, error)
	// GetUpdateStr return an update statement string that can later be used as a transaction's prepared statement
	//
	// Return error when:-
	// - Parameter is empty or missing
	GetUpdateStr(Fverinfo) (string, error)
	// GetDeleteStr return a delete statement string that can later be used as a transaction's prepared statement
	//
	// Return error when:-
	// - Parameter is empty or missing
	GetDeleteStr(int) (string, error)
}

//NewFverinfo return a new fverinfo struct based on parameters provided
func NewFverinfo(nodeID int, startDate string, endDate string, version string, verState int, remarks string) (Fverinfo, error) {
	if nodeID == 0 || startDate == "" || version == "" || verState == 0 {
		return Fverinfo{}, errors.New("New Fverinfo parameter cannot be empty")
	}
	return Fverinfo{
		NodeID:    nodeID,
		StartDate: startDate,
		EndDate:   endDate,
		Remarks:   remarks,
		Version:   version,
		VerState:  verState,
	}, nil
}
