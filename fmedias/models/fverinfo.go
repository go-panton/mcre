package models

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
	Insert(Fverinfo) error
	Update(Fverinfo) error
	Delete(int) error
	Find(int) (Fverinfo, error)
	GetInsertStr(Fverinfo) (string, error)
	GetUpdateStr(Fverinfo) (string, error)
	GetDeleteStr(int) (string, error)
}
