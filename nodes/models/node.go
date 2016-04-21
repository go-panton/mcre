package models

//Node struct define the data taht will be store or retrieve from the database
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
	// FindByDesc return the node based on the node description provided
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
	// - Node provided has missing field
	GetInsertStr(Node) (string, error)
	// GetUpdateStr return a sql string for update based on node provided
	//
	// Return error when:-
	// - Node provided has missing field
	GetUpdateStr(Node) (string, error)
	// GetDeleteStr return a sql string for delete based on node provided
	//
	// Return error when:-
	// - NodeID is 0
	GetDeleteStr(int) (string, error)
}
