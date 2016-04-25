package mysql

import (
	"fmt"
	"strconv"

	"database/sql"

	"errors"

	fmedias "github.com/go-panton/mcre/fmedias/models"
	nodes "github.com/go-panton/mcre/nodes/models"
	seqs "github.com/go-panton/mcre/seqs/models"
	users "github.com/go-panton/mcre/users/models"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sql.DB
}

type seqRepository struct {
	db *sql.DB
}

type nodeRepository struct {
	db *sqlx.DB
}

type fmediaRepository struct {
	db *sqlx.DB
}

type nodelinkRepository struct {
	db *sqlx.DB
}

type fverinfoRepository struct {
	db *sqlx.DB
}

//GetConnString returns connection string based on username, password and databaseName provided
func GetConnString(username, password, databaseName string) string {
	connString := username + ":" + password + "@/" + databaseName

	fmt.Println("Connection string: ", connString)

	return connString
}

//ConnectDatabase return a connection of database based on connString provided
func ConnectDatabase(connString string) *sql.DB {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Println("Error connecting to database")
	}
	return db
}

//NewUser return a UserRepository based on database connection provided
func NewUser(db *sql.DB) users.UserRepository {
	return &userRepository{db}
}

//NewSeq return a SeqRepository based on database connection provided
func NewSeq(db *sql.DB) seqs.SeqRepository {
	return &seqRepository{db}
}

//NewNode return a NodeRepository based on database connection provided
func NewNode(db *sqlx.DB) nodes.NodeRepository {
	return &nodeRepository{db}
}

//NewFmedia returns a FmediaRepository based on database connection provided
func NewFmedia(db *sqlx.DB) fmedias.FmediaRepository {
	return &fmediaRepository{db}
}

//NewNodeLink returns a NodelinkRepository based on database connection provided
func NewNodeLink(db *sqlx.DB) nodes.NodelinkRepository {
	return &nodelinkRepository{db}
}

//NewFverinfo returns a FverinfoRepository based on database connection provided
func NewFverinfo(db *sqlx.DB) fmedias.FverinfoRepository {
	return &fverinfoRepository{db}
}

func (r *userRepository) Insert(username, password string) error {
	insStat, err := r.db.Prepare("INSERT user SET username=?,password=?")

	if err != nil {
		return err
	}
	_, err1 := insStat.Exec(username, password)

	if err1 != nil {
		return err1
	}
	return nil
}

func (r *userRepository) Find(username string) (*users.User, error) {
	var resultName, resultPassword string
	err := r.db.QueryRow("SELECT * FROM user WHERE username=?", username).Scan(&resultName, &resultPassword)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	default:
		return &users.User{Username: resultName, Password: resultPassword}, nil
	}
}

func (r *userRepository) Verify(username, password string) (*users.User, error) {
	var resultName, resultPassword string
	err := r.db.QueryRow("SELECT * FROM user WHERE username=? AND password=?").Scan(&resultName, &resultPassword)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	default:
		return &users.User{Username: resultName, Password: password}, nil
	}
}

func (r *seqRepository) Find(key string) (int, error) {
	if key == "" {
		return 0, errors.New("The key is empty.")
	}

	var value int
	err := r.db.QueryRow("SELECT pseq FROM seqtbl WHERE PNAME=?", key).Scan(&value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (r *seqRepository) Update(key string, value int) error {
	if value < 1 {
		return errors.New("The update value should never be less than 1")
	}
	if key == "" {
		return errors.New("The key is empty.")
	}

	stmt, err := r.db.Prepare("UPDATE seqtbl SET pseq=? WHERE pname=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(value, key)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *nodeRepository) Insert(node nodes.Node) error {
	if node.NodeID == 0 || node.NodeDT == "" || node.NodeDesc == "" || node.NodeGID == 0 || node.NodeUID == 0 || node.NodeType == "" {
		return errors.New("Parameter cannot be empty")
	}

	stmt, err := r.db.Prepare("INSERT node SET nodeid=?,nodedesc=?,nodedt=?,nodegid=?,nodetype=?,nodeuid=?")
	if err != nil {
		return errors.New("Error on preparing insert statement: " + err.Error())
	}

	res, err := stmt.Exec(node.NodeID, node.NodeDesc, node.NodeDT, node.NodeGID, node.NodeType, node.NodeUID)
	if err != nil {
		return errors.New("Error executing insert statement: " + err.Error())
	}
	fmt.Println("Result of insert node statement: ", res)
	return nil
}

func (r *nodeRepository) Update(node nodes.Node) error {
	if node.NodeID == 0 || node.NodeDT == "" || node.NodeDesc == "" || node.NodeGID == 0 || node.NodeUID == 0 || node.NodeType == "" {
		return errors.New("Parameter cannot be empty")
	}

	stmt, err := r.db.Prepare("UPDATE node SET nodeDesc=?,nodegid=?,nodetype=?,nodeuid=? WHERE nodeid=?")
	if err != nil {
		return errors.New("Error preparing update statement" + err.Error())
	}

	res, err := stmt.Exec(node.NodeDesc, node.NodeGID, node.NodeType, node.NodeUID, node.NodeID)
	if err != nil {
		return errors.New("Error executing update statement: " + err.Error())
	}
	fmt.Println("Result of update node statement: ", res)

	return nil
}

func (r *nodeRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("Node ID cannot be 0")
	}

	stmt, err := r.db.Prepare("DELETE FROM node WHERE nodeid=?")
	if err != nil {
		return errors.New("Error preparing delete statement: " + err.Error())
	}

	res, err := stmt.Exec(nodeID)
	if err != nil {
		return errors.New("Error executing delete statement: " + err.Error())
	}
	fmt.Println("Result of delete node statement: ", res)

	return nil
}

func (r *nodeRepository) Find(nodeID int) (nodes.Node, error) {
	if nodeID == 0 {
		return nodes.Node{}, errors.New("Node ID cannot be 0")
	}

	var node nodes.Node
	err := r.db.QueryRowx("SELECT * FROM node WHERE nodeid=?", nodeID).StructScan(&node)
	if err != nil {
		if err == sql.ErrNoRows {
			return nodes.Node{}, errors.New("No result in database" + err.Error())
		}
		return nodes.Node{}, errors.New("Error querying database for result: " + err.Error())
	}

	return node, nil
}

func (r *nodeRepository) FindByDesc(nodeDesc string) ([]nodes.Node, error) {
	if nodeDesc == "" {
		return []nodes.Node{}, errors.New("Node Description cannot be empty")
	}

	node := []nodes.Node{}

	rows, err := r.db.Queryx("SELECT * FROM node where nodedesc=?", nodeDesc)
	if err != nil {
		if err == sql.ErrNoRows {
			return []nodes.Node{}, errors.New("No result in database: " + err.Error())
		}
		return []nodes.Node{}, errors.New("Error querying database: " + err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var r nodes.Node

		err = rows.StructScan(&r)
		if err != nil {
			return node, errors.New("Error while scanning result into Node")
		}
		node = append(node, r)
	}
	err = rows.Err()
	if err != nil {
		return []nodes.Node{}, errors.New("Error iterating rows" + err.Error())
	}
	return node, nil
}

func (r *nodeRepository) GetInsertStr(node nodes.Node) (insertStr string, err error) {
	if node.NodeID == 0 || node.NodeDT == "" || node.NodeDesc == "" || node.NodeGID == 0 {
		return "", errors.New("Parameter cannot be empty")
	}

	insertStr = "INSERT node SET nodeid=" + strconv.Itoa(node.NodeID) +
		",nodedesc=" + node.NodeDesc +
		",nodedt=" + node.NodeDT +
		",nodegid=" + strconv.Itoa(node.NodeGID) +
		",nodetype=" + node.NodeType +
		",nodeuid=" + strconv.Itoa(node.NodeUID)
	return insertStr, nil
}

func (r *nodeRepository) GetUpdateStr(node nodes.Node) (updateStr string, err error) {
	if node.NodeID == 0 || node.NodeDesc == "" || node.NodeGID == 0 || node.NodeUID == 0 || node.NodeType == "" {
		return "", errors.New("Parameter cannot be empty")
	}
	updateStr = "UPDATE node SET nodedesc=" + node.NodeDesc +
		",nodegid=" + strconv.Itoa(node.NodeGID) +
		",nodetype=" + node.NodeType +
		",nodeuid=" + strconv.Itoa(node.NodeUID) +
		" WHERE nodeid=" + strconv.Itoa(node.NodeID)

	return updateStr, nil
}

func (r *nodeRepository) GetDeleteStr(nodeID int) (deleteStr string, err error) {
	if nodeID == 0 {
		return "", errors.New("Node ID cannot be 0")
	}
	deleteStr = "DELETE FROM node WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}

func (r *fmediaRepository) Insert(fmedia fmedias.Fmedia) error {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return errors.New("Paramater cannot be empty")
	}

	stmt, err := r.db.Prepare("INSERT fmedia SET nodeid=?,fdesc=?,fext=?,ffulpath=?,fgname=?,foname=?,fremark=?,fsize=?,fstatus=?,ftype=?")
	if err != nil {
		return errors.New("Error preparing fmedia insert statement " + err.Error())
	}
	_, err = stmt.Exec(fmedia.NodeID, fmedia.FDesc, fmedia.FExt, fmedia.FFulPath, fmedia.FGName, fmedia.FOName, fmedia.FRemark, fmedia.FSize, fmedia.FStatus, fmedia.FType)
	if err != nil {
		return errors.New("Error executing fmedia insert statement " + err.Error())
	}
	return nil
}

func (r *fmediaRepository) Update(fmedia fmedias.Fmedia) error {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return errors.New("Paramater cannot be empty")
	}

	stmt, err := r.db.Prepare("UPDATE fmedia SET fdesc=?,fext=?,ffulpath=?,fgname=?,foname=?,fremark=?,fsize=?,fstatus=?,ftype=? WHERE nodeid=?")
	if err != nil {
		return errors.New("Error preparing fmedia update statement " + err.Error())
	}
	_, err = stmt.Exec(fmedia.FDesc, fmedia.FExt, fmedia.FFulPath, fmedia.FGName, fmedia.FOName, fmedia.FRemark, fmedia.FSize, fmedia.FStatus, fmedia.FType, fmedia.NodeID)
	if err != nil {
		return errors.New("Error executing fmedia update statement " + err.Error())
	}
	return nil
}

func (r *fmediaRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("Node ID cannot be 0 or empty")
	}
	stmt, err := r.db.Prepare("DELETE fmedia WHERE nodeid=?")
	if err != nil {
		return errors.New("Error preparing fmedia delete statement " + err.Error())
	}
	_, err = stmt.Exec(nodeID)
	if err != nil {
		return errors.New("Error executing fmedia delete statement " + err.Error())
	}
	return nil
}

func (r *fmediaRepository) Find(nodeID int) (fmedias.Fmedia, error) {
	if nodeID == 0 {
		return fmedias.Fmedia{}, errors.New("Node ID cannot be empty or 0")
	}

	var fmedia fmedias.Fmedia
	err := r.db.QueryRowx("SELECT * FROM fmedia WHERE nodeid=?", nodeID).StructScan(&fmedia)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmedias.Fmedia{}, errors.New("No results in database: " + err.Error())
		}
		return fmedias.Fmedia{}, errors.New("Error querying database: " + err.Error())
	}
	return fmedia, nil
}

func (r *fmediaRepository) FindByFileDesc(fileDesc string) ([]fmedias.Fmedia, error) {
	if fileDesc == "" {
		return []fmedias.Fmedia{}, errors.New("File description cannot be empty")
	}
	rows, err := r.db.Queryx("SELECT * FROM fmedia WHERE fdesc=?", fileDesc)
	if err != nil {
		if err == sql.ErrNoRows {
			return []fmedias.Fmedia{}, errors.New("No result from database: " + err.Error())
		}
		return []fmedias.Fmedia{}, errors.New("Error querying database: " + err.Error())
	}

	defer rows.Close()

	fmedia := []fmedias.Fmedia{}

	for rows.Next() {
		var r fmedias.Fmedia

		err := rows.StructScan(&r)
		if err != nil {
			return []fmedias.Fmedia{}, errors.New("Error Scanning result into fmedia struct: " + err.Error())
		}
		fmedia = append(fmedia, r)
	}
	err = rows.Err()
	if err != nil {
		return []fmedias.Fmedia{}, errors.New("Error while iterating through result: " + err.Error())
	}
	return fmedia, nil
}

func (r *fmediaRepository) GetInsertStr(fmedia fmedias.Fmedia) (string, error) {
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

func (r *fmediaRepository) GetUpdateStr(fmedia fmedias.Fmedia) (string, error) {
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

func (r *fmediaRepository) GetDeleteStr(nodeID int) (string, error) {
	if nodeID == 0 {
		return "", errors.New("Node ID cannot be 0 or empty")
	}
	deleteStr := "DELETE FROM fmedia WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}

func (r *nodelinkRepository) Insert(nl nodes.Nodelink) error {
	if nl.LinkCNodeID == 0 || nl.LinkPNodeID == 0 || nl.LinkType == "" {
		return errors.New("Parameter cannot be empty")
	}
	stmt, err := r.db.Prepare("INSERT nodelink SET linkpnodeid=?,linkcnodeid=?,linktype=?")
	if err != nil {
		return errors.New("Error preparing nodelink insert statement: " + err.Error())
	}
	_, err = stmt.Exec(nl.LinkPNodeID, nl.LinkCNodeID, nl.LinkType)
	if err != nil {
		return errors.New("Error executing nodelink insert statement: " + err.Error())
	}
	return nil
}

func (r *nodelinkRepository) Delete(childNodeID, parentNodeID int) error {
	if childNodeID == 0 || parentNodeID == 0 {
		return errors.New("Parameter cannot be empty")
	}
	stmt, err := r.db.Prepare("DELETE FROM nodelink WHERE linkpnodeid=?,linkcnodeid=?")
	if err != nil {
		return errors.New("Error preparing nodelink delete statement: " + err.Error())
	}
	_, err = stmt.Exec(parentNodeID, childNodeID)
	if err != nil {
		return errors.New("Error executing nodelink delete statement: " + err.Error())
	}
	return nil
}

func (r *nodelinkRepository) FindByChild(childNodeID int) ([]int, error) {
	if childNodeID == 0 {
		return []int{}, errors.New("Child Node ID cannot be empty or 0")
	}
	rows, err := r.db.Queryx("SELECT linkpnodeid FROM nodelink WHERE linkcnodeid=?", childNodeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []int{}, errors.New("No result from database: " + err.Error())
		}
		return []int{}, errors.New("Error querying database for result: " + err.Error())
	}
	defer rows.Close()

	parentNodeIDs := []int{}

	for rows.Next() {
		var r int

		err = rows.Scan(&r)
		if err != nil {
			return []int{}, errors.New("Error scanning result into parent nodeID slice: " + err.Error())
		}
		parentNodeIDs = append(parentNodeIDs, r)
	}
	err = rows.Err()
	if err != nil {
		return []int{}, errors.New("Error iterating rows: " + err.Error())
	}
	return parentNodeIDs, nil
}

func (r *nodelinkRepository) FindByParent(parentNodeID int) ([]int, error) {
	if parentNodeID == 0 {
		return []int{}, errors.New("Parent Node ID cannot be empty or 0")
	}
	rows, err := r.db.Queryx("SELECT linkcnodeid FROM nodelink WHERE linkpnodeid=?", parentNodeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []int{}, errors.New("No result from database: " + err.Error())
		}
		return []int{}, errors.New("Error querying database for result: " + err.Error())
	}
	defer rows.Close()

	childNodeIDs := []int{}

	for rows.Next() {
		var r int

		err = rows.Scan(&r)
		if err != nil {
			return []int{}, errors.New("Error scanning result into child nodeID slice: " + err.Error())
		}
		childNodeIDs = append(childNodeIDs, r)
	}
	err = rows.Err()
	if err != nil {
		return []int{}, errors.New("Error iterating rows: " + err.Error())
	}
	return childNodeIDs, nil
}

func (r *nodelinkRepository) FindExact(childNodeID int, parentNodeID int, linkType string) (nodes.Nodelink, error) {
	if childNodeID == 0 || parentNodeID == 0 || linkType == "" {
		return nodes.Nodelink{}, errors.New("Parameter cannot be empty")
	}

	var nl nodes.Nodelink

	err := r.db.QueryRowx("SELECT * FROM nodelink where linkpnodeid=?,linkcnodeid=?,linktype=?", parentNodeID, childNodeID, linkType).StructScan(&nl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nodes.Nodelink{}, errors.New("No result from database: " + err.Error())
		}
		return nodes.Nodelink{}, errors.New("Error scanning result into struct" + err.Error())
	}
	return nl, nil
}

func (r *nodelinkRepository) GetInsertStr(nl nodes.Nodelink) (string, error) {
	if nl.LinkCNodeID == 0 || nl.LinkPNodeID == 0 || nl.LinkType == "" {
		return "", errors.New("Parameter cannot be empty")
	}
	insertStr := "INSERT nodelink SET linkcnodeid=" + strconv.Itoa(nl.LinkCNodeID) +
		",linkpnodeid=" + strconv.Itoa(nl.LinkPNodeID) +
		",linkType=" + nl.LinkType

	return insertStr, nil
}

func (r *nodelinkRepository) GetDeleteStr(childNodeID, parentNodeID int) (string, error) {
	if childNodeID == 0 || parentNodeID == 0 {
		return "", errors.New("Parameter cannot be empty")
	}

	deleteStr := "DELETE FROM nodelink WHERE linkcnodeID=" + strconv.Itoa(childNodeID) +
		",linkpnodeid=" + strconv.Itoa(parentNodeID)

	return deleteStr, nil
}

func (r *fverinfoRepository) Insert(fverinfo fmedias.Fverinfo) error {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return errors.New("Parameter cannot be empty")
	}
	stmt, err := r.db.Prepare("INSERT fverinfo SET nodeid=?,enddate=?,remarks=?,startdate=?,version=?,verstate=?")
	if err != nil {
		return errors.New("Error preparing fverinfo insert statement: " + err.Error())
	}
	_, err = stmt.Exec(fverinfo.NodeID, fverinfo.EndDate, fverinfo.Remarks, fverinfo.StartDate, fverinfo.Version, fverinfo.VerState)
	if err != nil {
		return errors.New("Error executing fverinfo insert statement: " + err.Error())
	}
	return nil
}

func (r *fverinfoRepository) Update(fverinfo fmedias.Fverinfo) error {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return errors.New("Parameter cannot be empty")
	}
	stmt, err := r.db.Prepare("UPDATE fverinfo SET enddate=?,remarks=?,startdate=?,version=?,verstate=? WHERE nodeid=?")
	if err != nil {
		return errors.New("Error preparing fverinfo update statement")
	}
	_, err = stmt.Exec(fverinfo.EndDate, fverinfo.Remarks, fverinfo.StartDate, fverinfo.Version, fverinfo.VerState, fverinfo.NodeID)
	if err != nil {
		return errors.New("Error executing fverinfo update statement")
	}
	return nil
}

func (r *fverinfoRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("Node ID cannot be empty or 0")
	}
	stmt, err := r.db.Prepare("DELETE FROM fverinfo WHERE nodeid=?")
	if err != nil {
		return errors.New("Error preparing fverinfo delete statement")
	}
	_, err = stmt.Exec(nodeID)
	if err != nil {
		return errors.New("Error executing fverinfo delete statement")
	}
	return nil
}

func (r *fverinfoRepository) Find(nodeID int) (fmedias.Fverinfo, error) {
	if nodeID == 0 {
		return fmedias.Fverinfo{}, errors.New("Node ID cannot be empty or 0")
	}
	var fverinfo fmedias.Fverinfo
	err := r.db.QueryRowx("SELECT * FROM fverinfo WHERE nodeid=?", nodeID).StructScan(&fverinfo)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmedias.Fverinfo{}, errors.New("No result from database: " + err.Error())
		}
		return fmedias.Fverinfo{}, errors.New("Error querying result from database: " + err.Error())
	}
	return fverinfo, nil
}

func (r *fverinfoRepository) GetInsertStr(fverinfo fmedias.Fverinfo) (string, error) {
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

func (r *fverinfoRepository) GetUpdateStr(fverinfo fmedias.Fverinfo) (string, error) {
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

func (r *fverinfoRepository) GetDeleteStr(nodeID int) (string, error) {
	if nodeID == 0 {
		return "", errors.New("Node ID cannot be empty or 0")
	}
	deleteStr := "DELETE fverinfo WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}
