package model

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

var insertString string
var updateString string
var deleteString string

type NodeRepository interface {

	//
	Create(Node) error
	Update(Node) error
	Delete(int) error
	Find(int) (Node, error)
	FindByDesc(string) ([]Node, error)
	ConstructInsertString(Node) (string, error)
	ConstructUpdateString(Node) (string, error)
	ConstructDeleteString(int) (string, error)
}
