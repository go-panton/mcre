package mysql

import (
	"errors"
	"strconv"

	nodes "github.com/go-panton/mcre/nodes/models"
	seqs "github.com/go-panton/mcre/seqs/models"
	users "github.com/go-panton/mcre/users/models"
)

type mockUserRepository struct {
	UserArray []users.User
}

type mockSeqRepository struct{}

type mockNodeRepository struct {
	node nodes.Node
}

//NewMockUserRepository return a mock UserRepository for testing purpose
func NewMockUserRepository() users.UserRepository {
	return &mockUserRepository{}
}

//NewMockSeqRepository return a mock SeqRepository for testing purpose
func NewMockSeqRepository() seqs.SeqRepository {
	return &mockSeqRepository{}
}

//NewMockNodeRepository return a mock NodeRepository for testing purpose
func NewMockNodeRepository() nodes.NodeRepository {
	return &mockNodeRepository{}
}

func (m *mockUserRepository) Find(username string) (*users.User, error) {
	if username == "alex" {
		return &users.User{Username: "alex", Password: "root"}, nil
	}
	return nil, nil
}

func (m *mockUserRepository) Insert(username, password string) error {
	return nil
}

func (m *mockUserRepository) Verify(username, password string) (*users.User, error) {
	if username == "alex" && password == "root" {
		return &users.User{Username: "alex", Password: "root"}, nil
	}
	return nil, errors.New("Invalid User")
}

func (m *mockSeqRepository) Find(query string) (int, error) {
	if query == "" {
		return 0, errors.New("The query string is empty.")
	} else if query == "NODE" {
		return 200899, nil
	} else if query == "FILENAME" {
		return 761, nil
	}
	return 0, nil
}

func (m *mockSeqRepository) Update(query string, value int) error {
	if query == "" {
		return errors.New("The query string is empty")
	}
	if value < 1 {
		return errors.New("The value should not be less than 1")
	}
	return nil
}

func (m *mockNodeRepository) Insert(node nodes.Node) error {
	if node.NodeID == 0 || node.NodeDT == "" || node.NodeDesc == "" || node.NodeGID == 0 {
		return errors.New("Parameter cannot be empty")
	}

	return nil
}

func (m *mockNodeRepository) Update(node nodes.Node) error {
	if node.NodeID == 0 || node.NodeDT == "" || node.NodeDesc == "" || node.NodeGID == 0 {
		return errors.New("Parameter cannot be empty")
	}
	return nil
}

func (m *mockNodeRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("Node ID cannot be 0")
	}

	return nil
}

func (m *mockNodeRepository) Find(nodeID int) (nodes.Node, error) {
	if nodeID == 0 {
		return nodes.Node{}, errors.New("Node ID cannot be 0")
	}

	return nodes.Node{}, nil
}

func (m *mockNodeRepository) FindByDesc(nodeDesc string) ([]nodes.Node, error) {
	if nodeDesc == "" {
		return []nodes.Node{}, errors.New("Node Description cannot be empty")
	}

	return []nodes.Node{}, nil
}

func (m *mockNodeRepository) GetInsertStr(node nodes.Node) (string, error) {
	if node.NodeID == 0 || node.NodeDT == "" || node.NodeDesc == "" || node.NodeGID == 0 || node.NodeUID == 0 || node.NodeType == "" {
		return "", errors.New("Parameter cannot be empty")
	}
	insertStr := "INSERT node SET nodeid=" + strconv.Itoa(node.NodeID) +
		",nodedesc=" + node.NodeDesc +
		",nodedt=" + node.NodeDT +
		",nodegid=" + strconv.Itoa(node.NodeGID) +
		",nodetype=" + node.NodeType +
		",nodeuid=" + strconv.Itoa(node.NodeUID)

	return insertStr, nil
}

func (m *mockNodeRepository) GetUpdateStr(node nodes.Node) (string, error) {
	if node.NodeID == 0 || node.NodeDesc == "" || node.NodeGID == 0 || node.NodeUID == 0 {
		return "", errors.New("Parameter cannot be empty")
	}
	updateStr := "UPDATE node SET nodedesc=" + node.NodeDesc +
		",nodegid=" + strconv.Itoa(node.NodeGID) +
		",nodetype=" + node.NodeType +
		",nodeuid=" + strconv.Itoa(node.NodeUID) +
		" WHERE nodeid=" + strconv.Itoa(node.NodeID)

	return updateStr, nil
}

func (m *mockNodeRepository) GetDeleteStr(nodeID int) (string, error) {
	if nodeID == 0 {
		return "", errors.New("Node ID cannot be 0")
	}

	deleteStr := "DELETE FROM node WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}
