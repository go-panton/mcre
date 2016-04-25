package models

//Fmedia struct define the data that stores in the database
type Fmedia struct {
	NodeID   int
	FDesc    string
	FExt     string
	FFulPath string
	FGName   string
	FOName   string
	FRemark  string
	FSize    int
	FStatus  int
	FType    int
}

//FmediaRepository define the interface method
type FmediaRepository interface {
	// Insert insert the fmedia into database based on the fmedia provided
	//
	// Return error when:-
	// - insert failed
	// - Failed preparing statement
	Insert(Fmedia) error
	// Update update the value of fmedia in the database based on the fmedia provided
	//
	// Return error when:-
	// - update failed
	// - Failed preparing statement
	Update(Fmedia) error
	// Delete delete the fmedia from database based on the node id provided
	//
	// Return error when:-
	// - delete failed
	// - Failed preparing statement
	Delete(int) error
	// Find return the fmedia based on the node id provided
	//
	// Return error when:-
	// - nodeid is 0
	// - find failed
	// - no result
	// - invalid query
	Find(int) (Fmedia, error)
	// FindByFileDesc return the slice of Fmedia based on the file description provided
	//
	// Return error when:-
	// - file description is empty string
	// - find failed
	// - no result
	// - invalid query
	FindByFileDesc(string) ([]Fmedia, error)
	// GetInsertStr return a sql string for insert based on fmedia provided
	//
	// Return error when:-
	// - fmedia provided has missing field that are non nullable
	GetInsertStr(Fmedia) (string, error)
	// GetUpdateStr return a sql string for update based on fmedia provided
	//
	// Return error when:-
	// - Fmedia provided has missing field that are non nullable
	GetUpdateStr(Fmedia) (string, error)
	// GetDeleteStr return a sql string for delete based on fmedia provided
	//
	// Return error when:-
	// - NodeID is 0
	GetDeleteStr(int) (string, error)
}
