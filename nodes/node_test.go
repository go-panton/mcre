package nodes

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-panton/mcre/infra/store/mysql"
	"github.com/go-panton/mcre/nodes/models"
)

func TestInsert(t *testing.T) {
	nodes := mysql.NewMockNodeRepository()

	tests := []struct {
		Node models.Node
		Want error
	}{
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1}, nil},
		{models.Node{NodeID: 33333, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 242, NodeType: "FILE"}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeDesc: "testName1", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 242, NodeType: "FILE", NodeUID: 3212}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeUID: 1}, errors.New("Parameter cannot be empty")},
	}

	for _, test := range tests {
		err := nodes.Insert(test.Node)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestUpdate(t *testing.T) {
	nodes := mysql.NewMockNodeRepository()

	tests := []struct {
		Node models.Node
		Want error
	}{
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1}, nil},
		{models.Node{NodeID: 33333, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 242, NodeType: "FILE"}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeDesc: "testName1", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 242, NodeType: "FILE", NodeUID: 3212}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeUID: 1}, errors.New("Parameter cannot be empty")},
	}

	for _, test := range tests {
		err := nodes.Update(test.Node)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestDelete(t *testing.T) {
	nodes := mysql.NewMockNodeRepository()

	tests := []struct {
		nodeID int
		Want   error
	}{
		{22232, nil},
		{0, errors.New("Node ID cannot be 0")},
	}

	for _, test := range tests {
		err := nodes.Delete(test.nodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestFind(t *testing.T) {
	nodes := mysql.NewMockNodeRepository()

	tests := []struct {
		nodeID int
		Want   error
	}{
		{22232, nil},
		{0, errors.New("Node ID cannot be 0")},
	}

	for _, test := range tests {
		node, err := nodes.Find(test.nodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
		fmt.Printf("Result Find: %#v \n", node)
	}
}

func TestFindByDesc(t *testing.T) {
	nodes := mysql.NewMockNodeRepository()

	tests := []struct {
		nodeDesc string
		Want     error
	}{
		{"Hello", nil},
		{"", errors.New("Node Description cannot be empty")},
	}

	for _, test := range tests {
		res, err := nodes.FindByDesc(test.nodeDesc)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v,Want: %v \n", err.Error(), test.Want.Error())
			}
		}
		for _, node := range res {
			fmt.Printf("Result slice: %#v", node)
		}
	}
}

func TestGetInsertStr(t *testing.T) {
	nodes := mysql.NewMockNodeRepository()

	tests := []struct {
		Node models.Node
		Want error
	}{
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1}, nil},
		{models.Node{NodeID: 33333, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 242, NodeType: "FILE"}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeDesc: "testName1", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 242, NodeType: "FILE", NodeUID: 3212}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeUID: 1}, errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		res, err := nodes.GetInsertStr(test.Node)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
		fmt.Printf("Result string Insert: %v \n", res)
	}
}

func TestGetUpdateStr(t *testing.T) {
	nodes := mysql.NewMockNodeRepository()

	tests := []struct {
		Node models.Node
		Want error
	}{
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1}, nil},
		{models.Node{NodeID: 33333, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 242, NodeType: "FILE"}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeDesc: "testName1", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 242, NodeType: "FILE", NodeUID: 3212}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeGID: 424, NodeType: "FILE", NodeUID: 1}, errors.New("Parameter cannot be empty")},
		{models.Node{NodeID: 23232, NodeDesc: "testName", NodeDT: time.Now().Format("2006-01-02"), NodeGID: 424, NodeUID: 1}, errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		res, err := nodes.GetUpdateStr(test.Node)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
		fmt.Printf("Result string Update: %v \n", res)
	}
}

func TestGetDeleteStr(t *testing.T) {
	nodes := mysql.NewMockNodeRepository()

	tests := []struct {
		nodeID int
		Want   error
	}{
		{22232, nil},
		{0, errors.New("Node ID cannot be 0")},
	}
	for _, test := range tests {
		res, err := nodes.GetDeleteStr(test.nodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		}
		fmt.Printf("Result string Delete: %v \n", res)
	}
}
