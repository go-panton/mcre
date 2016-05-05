package models

import "errors"

//Nodelink struct defines the data that get stored in database
type Nodelink struct {
	LinkCNodeID int
	LinkPNodeID int
	LinkType    string
}

//NodelinkRepository defines the method interface to access nodelink struct
type NodelinkRepository interface {
	// Insert insert the nodelink into database based on the nodelink provided
	//
	// Return error when:-
	// - insert failed
	// - Failed preparing statement
	Insert(Nodelink) error
	// Delete delete the nodelink from database based on the child nodeID and parent nodeID provided
	//
	// Return error when:-
	// - child nodeID is 0 or parent nodeID is 0
	// - delete failed
	// - Failed preparing statement
	Delete(int, int) error
	// FindByChild return the slice of parent nodeID based on the child nodeID provided
	//
	// Return error when:-
	// - child nodeID is 0
	// - find failed
	// - no result
	// - invalid query
	FindByChild(int) ([]int, error)
	// FindByParent return the slice of child nodeID based on the parent nodeID provided
	//
	// Return error when:-
	// - parent nodeID is 0
	// - find failed
	// - no result
	// - invalid query
	FindByParent(int) ([]int, error)
	// FindExact return the nodelink based on the child nodeID and parent nodeID provided
	//
	// Return error when:-
	// - child nodeID or parent nodeID is 0 or linkType is empty string
	// - find failed
	// - no result
	// - invalid query
	FindExact(int, int, string) (Nodelink, error)
	// GetInsertStr return an insert statement that can later be used as a transaction's prepared statement based on parameters provided
	//
	// Return error when:-
	// - Parameters are missing or empty
	GetInsertStr(Nodelink) (string, error)
	// GetDeleteStr return a delete statement that can later be used as a transaction's prepared statement based on childNodeID and parentNodeID provided
	//
	// Return error when:-
	// - Parameters are missing or empty
	GetDeleteStr(int, int) (string, error)
}

//NewNodelink return a new Nodelink struct based on paramter input
func NewNodelink(childNodeID, parentNodeID int, linkType string) (Nodelink, error) {
	if childNodeID == 0 || parentNodeID == 0 || linkType == "" {
		return Nodelink{}, errors.New("Parameter cannot be empty or 0")
	}

	return Nodelink{
		LinkCNodeID: childNodeID,
		LinkPNodeID: parentNodeID,
		LinkType:    linkType,
	}, nil
}
