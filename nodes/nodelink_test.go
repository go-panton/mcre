package nodes

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-panton/mcre/infra/store/mysql"
	"github.com/go-panton/mcre/nodes/models"
)

func TestNodelinkInsert(t *testing.T) {
	nodelinks := mysql.NewMockNodelinkRepository()

	tests := []struct {
		Nodelink models.Nodelink
		Want     error
	}{
		{models.Nodelink{LinkCNodeID: 242422, LinkPNodeID: 434342, LinkType: "FILE"}, nil},
		{models.Nodelink{LinkPNodeID: 434342, LinkType: "FILE"}, errors.New("Parameter cannot be empty")},
		{models.Nodelink{LinkCNodeID: 242422, LinkType: "FILE"}, errors.New("Parameter cannot be empty")},
		{models.Nodelink{LinkCNodeID: 242422, LinkPNodeID: 434342}, errors.New("Parameter cannot be empty")},
	}

	for _, test := range tests {
		err := nodelinks.Insert(test.Nodelink)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v , Want %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestNodelinkDelete(t *testing.T) {
	nodelinks := mysql.NewMockNodelinkRepository()

	tests := []struct {
		parentNodeID int
		childNodeID  int
		Want         error
	}{
		{23232, 444544, nil},
		{0, 444544, errors.New("Parameter cannot be empty")},
		{23232, 0, errors.New("Parameter cannot be empty")},
	}

	for _, test := range tests {
		err := nodelinks.Delete(test.childNodeID, test.parentNodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v , Want: %v", err.Error(), test.Want.Error())
			}
		}
	}
}

func TestNodelinkFindByChild(t *testing.T) {
	nodelinks := mysql.NewMockNodelinkRepository()

	tests := []struct {
		NodeID int
		Want   error
	}{
		{224242, nil},
		{0, errors.New("Child Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		nodelink, err := nodelinks.FindByChild(test.NodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Nodelink: ", nodelink)
		}
	}
}

func TestNodelinkFindByParent(t *testing.T) {
	nodelinks := mysql.NewMockNodelinkRepository()

	tests := []struct {
		NodeID int
		Want   error
	}{
		{224242, nil},
		{0, errors.New("Parent Node ID cannot be empty or 0")},
	}
	for _, test := range tests {
		nodelink, err := nodelinks.FindByParent(test.NodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Nodelink: ", nodelink)
		}
	}
}

func TestFindExact(t *testing.T) {
	nodelinks := mysql.NewMockNodelinkRepository()

	tests := []struct {
		ChildNodeID  int
		ParentNodeID int
		LinkType     string
		Want         error
	}{
		{24242, 55454, "FILE", nil},
		{0, 54535, "FILE", errors.New("Parameter cannot be empty")},
		{342342, 0, "FILE", errors.New("Parameter cannot be empty")},
		{243324, 345433, "", errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		nodelink, err := nodelinks.FindExact(test.ChildNodeID, test.ParentNodeID, test.LinkType)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got %v , Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Nodelink: ", nodelink)
		}
	}
}

func TestNodelinkGetInsertStr(t *testing.T) {
	nodelinks := mysql.NewMockNodelinkRepository()

	tests := []struct {
		nl   models.Nodelink
		Want error
	}{
		{models.Nodelink{LinkCNodeID: 242422, LinkPNodeID: 434342, LinkType: "FILE"}, nil},
		{models.Nodelink{LinkPNodeID: 434342, LinkType: "FILE"}, errors.New("Parameter cannot be empty")},
		{models.Nodelink{LinkCNodeID: 242422, LinkType: "FILE"}, errors.New("Parameter cannot be empty")},
		{models.Nodelink{LinkCNodeID: 242422, LinkPNodeID: 434342}, errors.New("Parameter cannot be empty")},
	}
	for _, test := range tests {
		insertStr, err := nodelinks.GetInsertStr(test.nl)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v , Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Nodelink Insert String: ", insertStr)
		}
	}
}

func TestNodelinkGetDeleteStr(t *testing.T) {
	nodelinks := mysql.NewMockNodelinkRepository()

	tests := []struct {
		ChildNodeID  int
		ParentNodeID int
		Want         error
	}{
		{23232, 444544, nil},
		{0, 444544, errors.New("Parameter cannot be empty")},
		{23232, 0, errors.New("Parameter cannot be empty")},
	}

	for _, test := range tests {
		deleteStr, err := nodelinks.GetDeleteStr(test.ChildNodeID, test.ParentNodeID)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v , Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result Nodelink Delete String: " + deleteStr)
		}
	}
}

func TestNewNodeLink(t *testing.T) {
	tests := []struct {
		ChildNodeID  int
		ParentNodeID int
		LinkType     string
		Want         error
	}{
		{242424, 323232, "FILE", nil},
		{0, 323232, "FILE", errors.New("Parameter cannot be empty or 0")},
		{242424, 0, "FILE", errors.New("Parameter cannot be empty or 0")},
		{242424, 323232, "", errors.New("Parameter cannot be empty or 0")},
		{0, 0, "", errors.New("Parameter cannot be empty or 0")},
	}
	for _, test := range tests {
		nl, err := models.NewNodelink(test.ChildNodeID, test.ParentNodeID, test.LinkType)
		if err != nil {
			if err.Error() != test.Want.Error() {
				t.Errorf("Got: %v, Want: %v", err.Error(), test.Want.Error())
			}
		} else {
			fmt.Println("Result New NL: ", nl)
		}
	}
}
