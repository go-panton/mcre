package mysql

import (
	"errors"
	"strconv"

	fmedias "github.com/go-panton/mcre/fmedias/models"
	nodes "github.com/go-panton/mcre/nodes/models"
	seqs "github.com/go-panton/mcre/seqs/models"
	users "github.com/go-panton/mcre/users/models"
)

type mockUserRepository struct{}

type mockSeqRepository struct{}

type mockNodeRepository struct{}

type mockFmediaRepository struct{}

type mockFverinfoRepository struct{}

type mockNodelinkRepository struct{}

type mockConvqueueRepository struct{}

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

//NewMockFmediaRepository return a mock FmediaRepository for testing purpose
func NewMockFmediaRepository() fmedias.FmediaRepository {
	return &mockFmediaRepository{}
}

//NewMockFverinfoRepository return a mock FverinfoRepository for testing purpose
func NewMockFverinfoRepository() fmedias.FverinfoRepository {
	return &mockFverinfoRepository{}
}

//NewMockNodelinkRepository return a mock NodelinkRepository for testing purpose
func NewMockNodelinkRepository() nodes.NodelinkRepository {
	return &mockNodelinkRepository{}
}

//NewMockConvqueueRepository return a mock ConvqueueRepository for testing purpose
func NewMockConvqueueRepository() fmedias.ConvqueueRepository {
	return &mockConvqueueRepository{}
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
	if node.NodeID == 0 || node.NodeDesc == "" || node.NodeGID == 0 || node.NodeUID == 0 || node.NodeType == "" {
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

func (m *mockFmediaRepository) Insert(fmedia fmedias.Fmedia) error {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return errors.New("Paramater cannot be empty")
	}
	return nil
}

func (m *mockFmediaRepository) Update(fmedia fmedias.Fmedia) error {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return errors.New("Paramater cannot be empty")
	}
	return nil
}

func (m *mockFmediaRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("Node ID cannot be empty or 0")
	}
	return nil
}
func (m *mockFmediaRepository) Find(nodeID int) (fmedias.Fmedia, error) {
	if nodeID == 0 {
		return fmedias.Fmedia{}, errors.New("Node ID cannot be empty or 0")
	}
	return fmedias.Fmedia{}, nil
}

func (m *mockFmediaRepository) FindByFileDesc(fileDesc string) ([]fmedias.Fmedia, error) {
	if fileDesc == "" {
		return []fmedias.Fmedia{}, errors.New("File description cannot be empty")
	}
	return []fmedias.Fmedia{}, nil
}

func (m *mockFmediaRepository) GetInsertStr(fmedia fmedias.Fmedia) (string, error) {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return "", errors.New("Paramater cannot be empty")
	}
	insertStr := "INSERT fmedia SET nodeid=" + strconv.Itoa(fmedia.NodeID) +
		",fdesc=" + fmedia.FDesc +
		",fext=" + fmedia.FExt +
		",ffulpath=" + fmedia.FFulPath +
		",fgname=" + fmedia.FGName +
		",foname=" + fmedia.FOName +
		",fremark=" + fmedia.FRemark +
		",fsize=" + strconv.Itoa(fmedia.FSize) +
		",fstatus=" + strconv.Itoa(fmedia.FStatus) +
		",ftype=" + strconv.Itoa(fmedia.FType)

	return insertStr, nil
}

func (m *mockFmediaRepository) GetUpdateStr(fmedia fmedias.Fmedia) (string, error) {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return "", errors.New("Paramater cannot be empty")
	}
	updateStr := "UPDATE fmedia SET fdesc=" + fmedia.FDesc +
		",fext=" + fmedia.FExt +
		",ffulpath=" + fmedia.FFulPath +
		",fgname=" + fmedia.FGName +
		",foname=" + fmedia.FOName +
		",fremark=" + fmedia.FRemark +
		",fsize=" + strconv.Itoa(fmedia.FSize) +
		",fstatus=" + strconv.Itoa(fmedia.FStatus) +
		",ftype=" + strconv.Itoa(fmedia.FType) +
		" WHERE nodeid = " + strconv.Itoa(fmedia.NodeID)

	return updateStr, nil
}

func (m *mockFmediaRepository) GetDeleteStr(nodeID int) (string, error) {
	if nodeID == 0 {
		return "", errors.New("Node ID cannot be empty or 0")
	}
	deleteStr := "DELETE FROM fmedia WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}

func (m *mockFmediaRepository) CreateTx(nodeStr, fmStr, nlStr, fvStr, convStr string) error {
	if nodeStr == "" || fmStr == "" || nlStr == "" || fvStr == "" || convStr == "" {
		return errors.New("Parameter cannot be empty")
	}
	return nil
}

func (m *mockFverinfoRepository) Insert(fverinfo fmedias.Fverinfo) error {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return errors.New("Parameter cannot be empty")
	}
	return nil
}

func (m *mockFverinfoRepository) Update(fverinfo fmedias.Fverinfo) error {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return errors.New("Parameter cannot be empty")
	}
	return nil
}

func (m *mockFverinfoRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("Node ID cannot be empty or 0")
	}
	return nil
}

func (m *mockFverinfoRepository) Find(nodeID int) (fmedias.Fverinfo, error) {
	if nodeID == 0 {
		return fmedias.Fverinfo{}, errors.New("Node ID cannot be empty or 0")
	}
	return fmedias.Fverinfo{}, nil
}

func (m *mockFverinfoRepository) GetInsertStr(fverinfo fmedias.Fverinfo) (string, error) {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return "", errors.New("Parameter cannot be empty")
	}
	insertStr := "INSERT fverinfo SET nodeid=" + strconv.Itoa(fverinfo.NodeID) +
		",enddate=" + fverinfo.EndDate +
		",remarks=" + fverinfo.Remarks +
		",startdate=" + fverinfo.StartDate +
		",version=" + fverinfo.Version +
		",verstate=" + strconv.Itoa(fverinfo.VerState)

	return insertStr, nil
}

func (m *mockFverinfoRepository) GetUpdateStr(fverinfo fmedias.Fverinfo) (string, error) {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return "", errors.New("Parameter cannot be empty")
	}
	updateStr := "UPDATE fverinfo SET enddate=" + fverinfo.EndDate +
		",remarks=" + fverinfo.Remarks +
		",startdate=" + fverinfo.StartDate +
		",version=" + fverinfo.Version +
		",verstate=" + strconv.Itoa(fverinfo.VerState) +
		" WHERE nodeid=" + strconv.Itoa(fverinfo.NodeID)

	return updateStr, nil
}

func (m *mockFverinfoRepository) GetDeleteStr(nodeID int) (string, error) {
	if nodeID == 0 {
		return "", errors.New("Node ID cannot be empty or 0")
	}
	deleteStr := "DELETE FROM fverinfo WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}

func (m *mockNodelinkRepository) Insert(nl nodes.Nodelink) error {
	if nl.LinkCNodeID == 0 || nl.LinkPNodeID == 0 || nl.LinkType == "" {
		return errors.New("Parameter cannot be empty")
	}
	return nil
}

func (m *mockNodelinkRepository) Delete(childNodeID, parentNodeID int) error {
	if childNodeID == 0 || parentNodeID == 0 {
		return errors.New("Parameter cannot be empty")
	}
	return nil
}

func (m *mockNodelinkRepository) FindByChild(childNodeID int) ([]int, error) {
	if childNodeID == 0 {
		return []int{}, errors.New("Child Node ID cannot be empty or 0")
	}
	return []int{}, nil
}

func (m *mockNodelinkRepository) FindByParent(parentNodeID int) ([]int, error) {
	if parentNodeID == 0 {
		return []int{}, errors.New("Parent Node ID cannot be empty or 0")
	}
	return []int{}, nil
}

func (m *mockNodelinkRepository) FindExact(childNodeID int, parentNodeID int, linkType string) (nodes.Nodelink, error) {
	if childNodeID == 0 || parentNodeID == 0 || linkType == "" {
		return nodes.Nodelink{}, errors.New("Parameter cannot be empty")
	}
	return nodes.Nodelink{}, nil
}

func (m *mockNodelinkRepository) GetInsertStr(nl nodes.Nodelink) (string, error) {
	if nl.LinkCNodeID == 0 || nl.LinkPNodeID == 0 || nl.LinkType == "" {
		return "", errors.New("Parameter cannot be empty")
	}
	insertStr := "INSERT nodelink SET linkcnodeid=" + strconv.Itoa(nl.LinkCNodeID) +
		",linkpnodeid=" + strconv.Itoa(nl.LinkPNodeID) +
		",linktype=" + nl.LinkType

	return insertStr, nil
}

func (m *mockNodelinkRepository) GetDeleteStr(childNodeID, parentNodeID int) (string, error) {
	if childNodeID == 0 || parentNodeID == 0 {
		return "", errors.New("Parameter cannot be empty")
	}
	deleteStr := "DELETE FROM nodelink WHERE linkcnodeid=" + strconv.Itoa(childNodeID) +
		",linkpnodeid=" + strconv.Itoa(parentNodeID)

	return deleteStr, nil
}

func (m *mockConvqueueRepository) Insert(convqueue fmedias.Convqueue) error {
	if convqueue.NodeID == 0 || convqueue.Convtype == "" || convqueue.FExt == "" || convqueue.FFulpath == "" || convqueue.InsDate == "" || convqueue.Priority == 0 {
		return errors.New("Parameter cannot be empty")
	}
	return nil
}

func (m *mockConvqueueRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("Node ID cannot be empty or 0")
	}
	return nil
}

func (m *mockConvqueueRepository) GetInsertStr(convqueue fmedias.Convqueue) (string, error) {
	if convqueue.NodeID == 0 || convqueue.Convtype == "" || convqueue.FExt == "" || convqueue.FFulpath == "" || convqueue.InsDate == "" || convqueue.Priority == 0 {
		return "", errors.New("Parameter cannot be empty")
	}
	insertStr := "INSERT convqueue SET nodeid=" + strconv.Itoa(convqueue.NodeID) +
		",convtype=" + convqueue.Convtype +
		",fext=" + convqueue.FExt +
		",ffulpath=" + convqueue.FFulpath +
		",insdate=" + convqueue.InsDate +
		",priority=" + strconv.Itoa(convqueue.Priority)

	return insertStr, nil
}

func (m *mockConvqueueRepository) GetDeleteStr(nodeID int) (string, error) {
	if nodeID == 0 {
		return "", errors.New("Node ID cannot be empty or 0")
	}
	deleteStr := "DELETE FROM convqueue WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}
