package models

import "errors"

//Node struct define the data that will be store or retrieve from the database
type Node struct {
	NodeID    int
	NodeBits  int
	NodeDesc  string
	NodeDT    string //t.format("2006-01-02")
	NodeGID   int
	NodeHash  string
	NodeLevel int
	NodeType  string
	NodeUID   int
}

//NodeRepository define the interface method
type NodeRepository interface {

	// Insert insert the node into database based on the node provided
	//
	// Return error when:-
	// - insert failed
	// - Failed preparing statement
	Insert(Node) error
	// Update update the value of node in the database based on the node provided
	//
	// Return error when:-
	// - update failed
	// - Failed preparing statement
	Update(Node) error
	// Delete delete the node from database based on the id provided
	//
	// Return error when:-
	// - delete failed
	// - Failed preparing statement
	Delete(int) error
	// Find return the node based on the node id provided
	//
	// Return error when:-
	// - nodeid is 0
	// - find failed
	// - no result
	// - invalid query
	Find(int) (Node, error)
	// FindByDesc return slice of node based on the node description provided
	// Returning a slice of node because there may be multiple node having same description
	//
	// Return error when:-
	// - parameter is empty string
	// - find failed
	// - no result
	// - invalid query
	FindByDesc(string) ([]Node, error)
	// GetInsertStr return a sql string for insert based on node provided
	//
	// Return error when:-
	// - Node provided has missing field that are non nullable
	GetInsertStr(Node) (string, error)
	// GetUpdateStr return a sql string for update based on node provided
	//
	// Return error when:-
	// - Node provided has missing field that are non nullable
	GetUpdateStr(Node) (string, error)
	// GetDeleteStr return a sql string for delete based on node provided
	//
	// Return error when:-
	// - NodeID is 0
	GetDeleteStr(int) (string, error)
}

//NewNode returns a new Node based on parameters passed in
func NewNode(nodeID int, fileName string, fileDT string) (Node, error) {
	if nodeID == 0 || fileName == "" || fileDT == "" {
		return Node{}, errors.New("Parameters cannot be empty")
	}

	return Node{
		NodeID:   nodeID,
		NodeDesc: fileName,
		NodeDT:   fileDT,
		NodeGID:  200001,
		NodeType: "FILE",
		NodeUID:  1,
	}, nil
}
