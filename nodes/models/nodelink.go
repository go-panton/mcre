package models

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

	GetInsertStr(Nodelink) (string, error)

	GetDeleteStr(int, int) (string, error)
}
