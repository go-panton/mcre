package mysql

import (
	"fmt"
	"strconv"
	"strings"

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
	db *sqlx.DB
}

type seqRepository struct {
	db *sqlx.DB
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

type fmpolicyRepository struct {
	db *sqlx.DB
}

type convqueueRepository struct {
	db *sqlx.DB
}

//GetConnString returns connection string based on username, password and databaseName provided
func GetConnString(username, password, databaseName string) string {
	connString := username + ":" + password + "@/" + databaseName

	fmt.Println("Connection string: ", connString)

	return connString
}

//ConnectDatabase return a connection of database based on connString provided
func ConnectDatabase(connString string) *sqlx.DB {
	db, err := sqlx.Open("mysql", connString)

	if err != nil {
		fmt.Println("Error connecting to database")
	}
	return db
}

//NewUser return a UserRepository based on database connection provided
func NewUser(db *sqlx.DB) users.UserRepository {
	return &userRepository{db}
}

//NewSeq return a SeqRepository based on database connection provided
func NewSeq(db *sqlx.DB) seqs.SeqRepository {
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

//NewFmpolicy returns a FmpolicyRepository based on database connection provided
func NewFmpolicy(db *sqlx.DB) fmedias.FmpolicyRepository {
	return &fmpolicyRepository{db}
}

//NewConvqueue returns a ConvqueueRepository based on database connection provided
func NewConvqueue(db *sqlx.DB) fmedias.ConvqueueRepository {
	return &convqueueRepository{db}
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
		return 0, errors.New("seq Find: The key is empty")
	}

	var value int
	err := r.db.QueryRow("SELECT pseq FROM seqtbl WHERE PNAME=?", key).Scan(&value)
	if err != nil {
		return 0, err
	}
	return value + 1, nil
}

func (r *seqRepository) Update(key string, value int) error {
	if value < 1 {
		return errors.New("seq Update: The update value should never be less than 1")
	}
	if key == "" {
		return errors.New("seq Update: The key is empty")
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
		return errors.New("node Insert: Parameter cannot be empty")
	}

	stmt, err := r.db.Prepare("INSERT node SET nodeid=?,nodedesc=?,nodedt=?,nodegid=?,nodetype=?,nodeuid=?")
	if err != nil {
		return errors.New("node Insert: Error on preparing insert statement: " + err.Error())
	}

	res, err := stmt.Exec(node.NodeID, node.NodeDesc, node.NodeDT, node.NodeGID, node.NodeType, node.NodeUID)
	if err != nil {
		return errors.New("node Insert: Error executing insert statement: " + err.Error())
	}
	fmt.Println("Result of insert node statement: ", res)
	return nil
}

func (r *nodeRepository) Update(node nodes.Node) error {
	if node.NodeID == 0 || node.NodeDT == "" || node.NodeDesc == "" || node.NodeGID == 0 || node.NodeUID == 0 || node.NodeType == "" {
		return errors.New("node Update: Parameter cannot be empty")
	}

	stmt, err := r.db.Prepare("UPDATE node SET nodeDesc=?,nodegid=?,nodetype=?,nodeuid=? WHERE nodeid=?")
	if err != nil {
		return errors.New("node Update: Error preparing update statement" + err.Error())
	}

	res, err := stmt.Exec(node.NodeDesc, node.NodeGID, node.NodeType, node.NodeUID, node.NodeID)
	if err != nil {
		return errors.New("node Update: Error executing update statement: " + err.Error())
	}
	fmt.Println("Result of update node statement: ", res)

	return nil
}

func (r *nodeRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("node Delete: Node ID cannot be 0")
	}

	stmt, err := r.db.Prepare("DELETE FROM node WHERE nodeid=?")
	if err != nil {
		return errors.New("node Delete: Error preparing delete statement: " + err.Error())
	}

	res, err := stmt.Exec(nodeID)
	if err != nil {
		return errors.New("node Delete: Error executing delete statement: " + err.Error())
	}
	fmt.Println("Result of delete node statement: ", res)

	return nil
}

func (r *nodeRepository) Find(nodeID int) (nodes.Node, error) {
	if nodeID == 0 {
		return nodes.Node{}, errors.New("node Find: Node ID cannot be 0")
	}

	var node nodes.Node
	err := r.db.QueryRowx("SELECT * FROM node WHERE nodeid=?", nodeID).StructScan(&node)
	if err != nil {
		if err == sql.ErrNoRows {
			return nodes.Node{}, errors.New("node Find: No result in database" + err.Error())
		}
		return nodes.Node{}, errors.New("node Find: Error querying database for result: " + err.Error())
	}

	return node, nil
}

func (r *nodeRepository) FindByDesc(nodeDesc string) ([]nodes.Node, error) {
	if nodeDesc == "" {
		return []nodes.Node{}, errors.New("node FindByDesc: Node Description cannot be empty")
	}

	node := []nodes.Node{}

	rows, err := r.db.Queryx("SELECT * FROM node where nodedesc=?", nodeDesc)
	if err != nil {
		if err == sql.ErrNoRows {
			return []nodes.Node{}, errors.New("node FindByDesc: No result in database: " + err.Error())
		}
		return []nodes.Node{}, errors.New("node FindByDesc: Error querying database: " + err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var res nodes.Node

		err = rows.StructScan(&res)
		if err != nil {
			return node, errors.New("node FindByDesc: Error while scanning result into Node")
		}
		node = append(node, res)
	}
	err = rows.Err()
	if err != nil {
		return []nodes.Node{}, errors.New("node FindByDesc: Error iterating rows" + err.Error())
	}
	return node, nil
}

func (r *nodeRepository) GetInsertStr(node nodes.Node) (insertStr string, err error) {
	if node.NodeID == 0 || node.NodeDT == "" || node.NodeDesc == "" || node.NodeGID == 0 {
		return "", errors.New("node GetInsertStr: Parameter cannot be empty")
	}

	insertStr = "INSERT node SET nodeid=" + strconv.Itoa(node.NodeID) +
		",nodedesc=\"" + node.NodeDesc +
		"\",nodedt=\"" + node.NodeDT +
		"\",nodegid=" + strconv.Itoa(node.NodeGID) +
		",nodetype=\"" + node.NodeType +
		"\",nodeuid=" + strconv.Itoa(node.NodeUID)
	return insertStr, nil
}

func (r *nodeRepository) GetUpdateStr(node nodes.Node) (updateStr string, err error) {
	if node.NodeID == 0 || node.NodeDesc == "" || node.NodeGID == 0 || node.NodeUID == 0 || node.NodeType == "" {
		return "", errors.New("node GetUpdateStr: Parameter cannot be empty")
	}
	updateStr = "UPDATE node SET nodedesc=\"" + node.NodeDesc +
		"\",nodegid=" + strconv.Itoa(node.NodeGID) +
		",nodetype=\"" + node.NodeType +
		"\",nodeuid=" + strconv.Itoa(node.NodeUID) +
		" WHERE nodeid=" + strconv.Itoa(node.NodeID)

	return updateStr, nil
}

func (r *nodeRepository) GetDeleteStr(nodeID int) (deleteStr string, err error) {
	if nodeID == 0 {
		return "", errors.New("node GetDeleteStr: Node ID cannot be 0")
	}
	deleteStr = "DELETE FROM node WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}

func (r *fmediaRepository) Insert(fmedia fmedias.Fmedia) error {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return errors.New("fmedia Insert: Paramater cannot be empty")
	}

	stmt, err := r.db.Prepare("INSERT fmedia SET nodeid=?,fdesc=?,fext=?,ffulpath=?,fgname=?,foname=?,fremark=?,fsize=?,fstatus=?,ftype=?")
	if err != nil {
		return errors.New("fmedia Insert: Error preparing fmedia insert statement " + err.Error())
	}
	_, err = stmt.Exec(fmedia.NodeID, fmedia.FDesc, fmedia.FExt, fmedia.FFulPath, fmedia.FGName, fmedia.FOName, fmedia.FRemark, fmedia.FSize, fmedia.FStatus, fmedia.FType)
	if err != nil {
		return errors.New("fmedia Insert: Error executing fmedia insert statement " + err.Error())
	}
	return nil
}

func (r *fmediaRepository) Update(fmedia fmedias.Fmedia) error {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return errors.New("fmedia Update: Paramater cannot be empty")
	}

	stmt, err := r.db.Prepare("UPDATE fmedia SET fdesc=?,fext=?,ffulpath=?,fgname=?,foname=?,fremark=?,fsize=?,fstatus=?,ftype=? WHERE nodeid=?")
	if err != nil {
		return errors.New("fmedia Update: Error preparing fmedia update statement " + err.Error())
	}
	_, err = stmt.Exec(fmedia.FDesc, fmedia.FExt, fmedia.FFulPath, fmedia.FGName, fmedia.FOName, fmedia.FRemark, fmedia.FSize, fmedia.FStatus, fmedia.FType, fmedia.NodeID)
	if err != nil {
		return errors.New("fmedia Update: Error executing fmedia update statement " + err.Error())
	}
	return nil
}

func (r *fmediaRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("fmedia Delete: Node ID cannot be 0 or empty")
	}
	stmt, err := r.db.Prepare("DELETE fmedia WHERE nodeid=?")
	if err != nil {
		return errors.New("fmedia Delete: Error preparing fmedia delete statement " + err.Error())
	}
	_, err = stmt.Exec(nodeID)
	if err != nil {
		return errors.New("fmedia Delete: Error executing fmedia delete statement " + err.Error())
	}
	return nil
}

func (r *fmediaRepository) Find(nodeID int) (fmedias.Fmedia, error) {
	if nodeID == 0 {
		return fmedias.Fmedia{}, errors.New("fmedia Find: Node ID cannot be empty or 0")
	}

	var fmedia fmedias.Fmedia
	err := r.db.QueryRowx("SELECT * FROM fmedia WHERE nodeid=?", nodeID).StructScan(&fmedia)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmedias.Fmedia{}, errors.New("fmedia Find: No results in database: " + err.Error())
		}
		return fmedias.Fmedia{}, errors.New("fmedia Find: Error querying database: " + err.Error())
	}
	return fmedia, nil
}

func (r *fmediaRepository) FindByFileDesc(fileDesc string) ([]fmedias.Fmedia, error) {
	if fileDesc == "" {
		return []fmedias.Fmedia{}, errors.New("fmedia FindByDesc: File description cannot be empty")
	}
	rows, err := r.db.Queryx("SELECT * FROM fmedia WHERE fdesc=?", fileDesc)
	if err != nil {
		if err == sql.ErrNoRows {
			return []fmedias.Fmedia{}, errors.New("fmedia FindByDesc: No result from database: " + err.Error())
		}
		return []fmedias.Fmedia{}, errors.New("fmedia FindByDesc: Error querying database: " + err.Error())
	}

	defer rows.Close()

	fmedia := []fmedias.Fmedia{}

	for rows.Next() {
		var res fmedias.Fmedia

		err := rows.StructScan(&res)
		if err != nil {
			return []fmedias.Fmedia{}, errors.New("fmedia FindByDesc: Error Scanning result into fmedia struct: " + err.Error())
		}
		fmedia = append(fmedia, res)
	}
	err = rows.Err()
	if err != nil {
		return []fmedias.Fmedia{}, errors.New("fmedia FindByDesc: Error while iterating through result: " + err.Error())
	}
	return fmedia, nil
}

func (r *fmediaRepository) GetInsertStr(fmedia fmedias.Fmedia) (string, error) {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return "", errors.New("fmedia GetInsertStr: Paramater cannot be empty")
	}

	//need to escape the backslash else insert will fail
	fmedia.FFulPath = strings.Replace(fmedia.FFulPath, "\\", "\\\\", -1)

	insertStr := "INSERT fmedia SET nodeid=" + strconv.Itoa(fmedia.NodeID) +
		",fdesc=\"" + fmedia.FDesc +
		"\",fext=\"" + fmedia.FExt +
		"\",ffulpath=\"" + fmedia.FFulPath +
		"\",fgname=\"" + fmedia.FGName +
		"\",foname=\"" + fmedia.FOName +
		"\",fremark=\"" + fmedia.FRemark +
		"\",fsize=" + strconv.Itoa(fmedia.FSize) +
		",fstatus=" + strconv.Itoa(fmedia.FStatus) +
		",ftype=" + strconv.Itoa(fmedia.FType)

	return insertStr, nil
}

func (r *fmediaRepository) GetUpdateStr(fmedia fmedias.Fmedia) (string, error) {
	if fmedia.NodeID == 0 || fmedia.FDesc == "" || fmedia.FExt == "" || fmedia.FFulPath == "" || fmedia.FGName == "" || fmedia.FOName == "" || fmedia.FSize == 0 {
		return "", errors.New("fmedia GetUpdateStr: Paramater cannot be empty")
	}

	fmedia.FFulPath = strings.Replace(fmedia.FFulPath, "\\", "\\\\", -1)

	updateStr := "UPDATE fmedia SET fdesc=\"" + fmedia.FDesc +
		"\",fext=\"" + fmedia.FExt +
		"\",ffulpath=\"" + fmedia.FFulPath +
		"\",fgname=\"" + fmedia.FGName +
		"\",foname=\"" + fmedia.FOName +
		"\",fremark=\"" + fmedia.FRemark +
		"\",fsize=" + strconv.Itoa(fmedia.FSize) +
		",fstatus=" + strconv.Itoa(fmedia.FStatus) +
		",ftype=" + strconv.Itoa(fmedia.FType) +
		" WHERE nodeid = " + strconv.Itoa(fmedia.NodeID)

	return updateStr, nil
}

func (r *fmediaRepository) GetDeleteStr(nodeID int) (string, error) {
	if nodeID == 0 {
		return "", errors.New("fmedia GetDeleteStr: Node ID cannot be 0 or empty")
	}
	deleteStr := "DELETE FROM fmedia WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}

func (r *nodelinkRepository) Insert(nl nodes.Nodelink) error {
	if nl.LinkCNodeID == 0 || nl.LinkPNodeID == 0 || nl.LinkType == "" {
		return errors.New("nodelink Insert: Parameter cannot be empty")
	}
	stmt, err := r.db.Prepare("INSERT nodelink SET linkpnodeid=?,linkcnodeid=?,linktype=?")
	if err != nil {
		return errors.New("nodelink Insert: Error preparing nodelink insert statement: " + err.Error())
	}
	_, err = stmt.Exec(nl.LinkPNodeID, nl.LinkCNodeID, nl.LinkType)
	if err != nil {
		return errors.New("nodelink Insert: Error executing nodelink insert statement: " + err.Error())
	}
	return nil
}

func (r *nodelinkRepository) Delete(childNodeID, parentNodeID int) error {
	if childNodeID == 0 || parentNodeID == 0 {
		return errors.New("nodelink Delete: Parameter cannot be empty")
	}
	stmt, err := r.db.Prepare("DELETE FROM nodelink WHERE linkpnodeid=? AND linkcnodeid=?")
	if err != nil {
		return errors.New("nodelink Delete: Error preparing nodelink delete statement: " + err.Error())
	}
	_, err = stmt.Exec(parentNodeID, childNodeID)
	if err != nil {
		return errors.New("nodelink Delete: Error executing nodelink delete statement: " + err.Error())
	}
	return nil
}

func (r *nodelinkRepository) FindByChild(childNodeID int) ([]int, error) {
	if childNodeID == 0 {
		return []int{}, errors.New("nodelink FindByChild: Child Node ID cannot be empty or 0")
	}
	rows, err := r.db.Queryx("SELECT linkpnodeid FROM nodelink WHERE linkcnodeid=?", childNodeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []int{}, errors.New("nodelink FindByChild: No result from database: " + err.Error())
		}
		return []int{}, errors.New("nodelink FindByChild: Error querying database for result: " + err.Error())
	}
	defer rows.Close()

	parentNodeIDs := []int{}

	for rows.Next() {
		var res int

		err = rows.Scan(&res)
		if err != nil {
			return []int{}, errors.New("nodelink FindByChild: Error scanning result into parent nodeID slice: " + err.Error())
		}
		parentNodeIDs = append(parentNodeIDs, res)
	}
	err = rows.Err()
	if err != nil {
		return []int{}, errors.New("nodelink FindByChild: Error iterating rows: " + err.Error())
	}
	return parentNodeIDs, nil
}

func (r *nodelinkRepository) FindByParent(parentNodeID int, linkType string) ([]int, error) {
	if parentNodeID == 0 {
		return []int{}, errors.New("nodelink FindByParent: Parameter cannot be empty or 0")
	}
	rows, err := r.db.Queryx("SELECT linkcnodeid FROM nodelink WHERE linkpnodeid=? AND linktype=?", parentNodeID, linkType)
	if err != nil {
		if err == sql.ErrNoRows {
			return []int{}, errors.New("nodelink FindByParent: No result from database: " + err.Error())
		}
		return []int{}, errors.New("nodelink FindByParent: Error querying database for result: " + err.Error())
	}
	defer rows.Close()

	childNodeIDs := []int{}

	for rows.Next() {
		var res int

		err = rows.Scan(&res)
		if err != nil {
			return []int{}, errors.New("nodelink FindByParent: Error scanning result into child nodeID slice: " + err.Error())
		}
		childNodeIDs = append(childNodeIDs, res)
	}
	err = rows.Err()
	if err != nil {
		return []int{}, errors.New("nodelink FindByParent: Error iterating rows: " + err.Error())
	}
	return childNodeIDs, nil
}

func (r *nodelinkRepository) FindExact(childNodeID int, parentNodeID int, linkType string) (nodes.Nodelink, error) {
	if childNodeID == 0 || parentNodeID == 0 || linkType == "" {
		return nodes.Nodelink{}, errors.New("nodelink FindExact: Parameter cannot be empty")
	}

	var nl nodes.Nodelink

	err := r.db.QueryRowx("SELECT * FROM nodelink where linkpnodeid=? AND linkcnodeid=? AND linktype=?", parentNodeID, childNodeID, linkType).StructScan(&nl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nodes.Nodelink{}, errors.New("nodelink FindExact: No result from database: " + err.Error())
		}
		return nodes.Nodelink{}, errors.New("nodelink FindExact: Error scanning result into struct" + err.Error())
	}
	return nl, nil
}

func (r *nodelinkRepository) GetInsertStr(nl nodes.Nodelink) (string, error) {
	if nl.LinkCNodeID == 0 || nl.LinkPNodeID == 0 || nl.LinkType == "" {
		return "", errors.New("nodelink GetInsertStr: Parameter cannot be empty")
	}
	insertStr := "INSERT nodelink SET linkcnodeid=" + strconv.Itoa(nl.LinkCNodeID) +
		",linkpnodeid=" + strconv.Itoa(nl.LinkPNodeID) +
		",linkType=\"" + nl.LinkType + "\""

	return insertStr, nil
}

func (r *nodelinkRepository) GetDeleteStr(childNodeID, parentNodeID int) (string, error) {
	if childNodeID == 0 || parentNodeID == 0 {
		return "", errors.New("nodelink GetDeleteStr: Parameter cannot be empty")
	}

	deleteStr := "DELETE FROM nodelink WHERE linkcnodeID=" + strconv.Itoa(childNodeID) +
		" AND linkpnodeid=" + strconv.Itoa(parentNodeID)

	return deleteStr, nil
}

func (r *fverinfoRepository) Insert(fverinfo fmedias.Fverinfo) error {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return errors.New("fverinfo Insert: Parameter cannot be empty")
	}
	stmt, err := r.db.Prepare("INSERT fverinfo SET nodeid=?,enddate=?,remarks=?,startdate=?,version=?,verstate=?")
	if err != nil {
		return errors.New("fverinfo Insert: Error preparing fverinfo insert statement: " + err.Error())
	}
	_, err = stmt.Exec(fverinfo.NodeID, fverinfo.EndDate, fverinfo.Remarks, fverinfo.StartDate, fverinfo.Version, fverinfo.VerState)
	if err != nil {
		return errors.New("fverinfo Insert: Error executing fverinfo insert statement: " + err.Error())
	}
	return nil
}

func (r *fverinfoRepository) Update(fverinfo fmedias.Fverinfo) error {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return errors.New("fverinfo Update: Parameter cannot be empty")
	}
	stmt, err := r.db.Prepare("UPDATE fverinfo SET enddate=?,remarks=?,startdate=?,version=?,verstate=? WHERE nodeid=?")
	if err != nil {
		return errors.New("fverinfo Update: Error preparing fverinfo update statement")
	}
	_, err = stmt.Exec(fverinfo.EndDate, fverinfo.Remarks, fverinfo.StartDate, fverinfo.Version, fverinfo.VerState, fverinfo.NodeID)
	if err != nil {
		return errors.New("fverinfo Update: Error executing fverinfo update statement")
	}
	return nil
}

func (r *fverinfoRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("fverinfo Delete: Node ID cannot be empty or 0")
	}
	stmt, err := r.db.Prepare("DELETE FROM fverinfo WHERE nodeid=?")
	if err != nil {
		return errors.New("fverinfo Delete: Error preparing fverinfo delete statement")
	}
	_, err = stmt.Exec(nodeID)
	if err != nil {
		return errors.New("fverinfo Delete: Error executing fverinfo delete statement")
	}
	return nil
}

func (r *fverinfoRepository) Find(nodeID int) (fmedias.Fverinfo, error) {
	if nodeID == 0 {
		return fmedias.Fverinfo{}, errors.New("fverinfo Find: Node ID cannot be empty or 0")
	}
	var fverinfo fmedias.Fverinfo
	err := r.db.QueryRowx("SELECT * FROM fverinfo WHERE nodeid=?", nodeID).StructScan(&fverinfo)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmedias.Fverinfo{}, errors.New("fverinfo Find: No result from database: " + err.Error())
		}
		return fmedias.Fverinfo{}, errors.New("fverinfo Find: Error querying result from database: " + err.Error())
	}
	return fverinfo, nil
}

func (r *fverinfoRepository) GetInsertStr(fverinfo fmedias.Fverinfo) (string, error) {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return "", errors.New("fverinfo GetInsertStr: Parameter cannot be empty")
	}
	insertStr := "INSERT fverinfo SET nodeid=" + strconv.Itoa(fverinfo.NodeID) +
		",enddate=\"" + fverinfo.EndDate +
		"\",remarks=\"" + fverinfo.Remarks +
		"\",startdate=\"" + fverinfo.StartDate +
		"\",version=\"" + fverinfo.Version +
		"\",verstate=" + strconv.Itoa(fverinfo.VerState)

	return insertStr, nil
}

func (r *fverinfoRepository) GetUpdateStr(fverinfo fmedias.Fverinfo) (string, error) {
	if fverinfo.NodeID == 0 || fverinfo.StartDate == "" || fverinfo.Version == "" || fverinfo.VerState == 0 {
		return "", errors.New("fverinfo GetUpdateStr: Parameter cannot be empty")
	}
	updateStr := "UPDATE fverinfo SET enddate=\"" + fverinfo.EndDate +
		"\",remarks=\"" + fverinfo.Remarks +
		"\",startdate=\"" + fverinfo.StartDate +
		"\",version=\"" + fverinfo.Version +
		"\",verstate=" + strconv.Itoa(fverinfo.VerState) +
		" WHERE nodeid=" + strconv.Itoa(fverinfo.NodeID)

	return updateStr, nil
}

func (r *fverinfoRepository) GetDeleteStr(nodeID int) (string, error) {
	if nodeID == 0 {
		return "", errors.New("fverinfo GetDeleteStr: Node ID cannot be empty or 0")
	}
	deleteStr := "DELETE fverinfo WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}

func (r *fmpolicyRepository) Insert(fmpolicy fmedias.Fmpolicy) error {
	if fmpolicy.NodeID == 0 {
		return errors.New("fmpolicy Insert: NodeID cannot be empty or 0")
	}
	stmt, err := r.db.Prepare("INSERT fmpolicy SET fmpdownload=?, fmprevise=?, fmpview=?, fmpugid=?, fmpugtype=?, nodeid=?")
	if err != nil {
		return errors.New("fmpolicy Insert: Error when preparing fmpolicy prepared statement: " + err.Error())
	}
	_, err = stmt.Exec(fmpolicy.FmpDownload, fmpolicy.FmpRevise, fmpolicy.FmpView, fmpolicy.FmpUGID, fmpolicy.FmpUGType, fmpolicy.NodeID)
	if err != nil {
		return errors.New("fmpolicy Insert: Error when executing fmpolicy insert statement: " + err.Error())
	}
	return nil
}

func (r *fmpolicyRepository) Update(fmpolicy fmedias.Fmpolicy) error {
	if fmpolicy.NodeID == 0 || fmpolicy.FmpID == 0 {
		return errors.New("fmpolicy Update: Parameter cannot be empty or 0")
	}
	stmt, err := r.db.Prepare("UPDATE fmpolicy SET fmpdownload=?, fmprevise=?, fmpview=?, fmpugid=?, fmpugtype=? WHERE fmpid=? AND nodeid=?")
	if err != nil {
		return errors.New("fmpolicy Update: Error when preparing fmpolicy prepared statement: " + err.Error())
	}
	_, err = stmt.Exec(fmpolicy.FmpDownload, fmpolicy.FmpRevise, fmpolicy.FmpView, fmpolicy.FmpUGID, fmpolicy.FmpUGType, fmpolicy.FmpID, fmpolicy.NodeID)
	if err != nil {
		return errors.New("fmpolicy Update: Error when executing fmpolicy statement: " + err.Error())
	}
	return nil
}

func (r *fmpolicyRepository) Delete(fmpID int) error {
	if fmpID == 0 {
		return errors.New("fmpolicy Delete: Fmp ID cannot be empty or 0")
	}
	stmt, err := r.db.Prepare("DELETE FROM fmpolicy WHERE fmpid=?")
	if err != nil {
		return errors.New("fmpolicy Delete: Error when preparing fmpolicy delete statement: " + err.Error())
	}
	_, err = stmt.Exec(fmpID)
	if err != nil {
		return errors.New("fmpolicy Delete: Error when executing fmpolicy delete statement: " + err.Error())
	}
	return nil
}

func (r *fmpolicyRepository) Find(fmpID int) (fmedias.Fmpolicy, error) {
	if fmpID == 0 {
		return fmedias.Fmpolicy{}, errors.New("fmpolicy Find: Fmp ID cannot be empty or 0")
	}
	var fmpolicy fmedias.Fmpolicy
	err := r.db.QueryRowx("SELECT * FROM fmpolicy WHERE fmpid=?", fmpID).StructScan(&fmpolicy)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmedias.Fmpolicy{}, errors.New("fmpolicy Find: No result from database: " + err.Error())
		}
		return fmedias.Fmpolicy{}, errors.New("fmpolicy Find: Error querying results from database: " + err.Error())
	}
	return fmpolicy, nil
}

func (r *fmpolicyRepository) FindUsingNodeID(nodeID int) ([]fmedias.Fmpolicy, error) {
	if nodeID == 0 {
		return []fmedias.Fmpolicy{}, errors.New("fmpolicy FindUsingNodeID: Node ID cannot be empty or 0")
	}

	rows, err := r.db.Queryx("SELECT * FROM fmpolicy WHERE nodeid=?", nodeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []fmedias.Fmpolicy{}, errors.New("fmpolicy FindUsingNodeID: No result from database: " + err.Error())
		}
		return []fmedias.Fmpolicy{}, errors.New("fmpolicy FindUsingNodeID: Error querying database: " + err.Error())
	}
	defer rows.Close()

	var fmps []fmedias.Fmpolicy

	for rows.Next() {
		var res fmedias.Fmpolicy
		err := rows.StructScan(&res)
		if err != nil {
			return []fmedias.Fmpolicy{}, errors.New("fmpolicy FindUsingNodeID: Error scanning result into struct: " + err.Error())
		}
		fmps = append(fmps, res)
	}
	err = rows.Err()
	if err != nil {
		return []fmedias.Fmpolicy{}, errors.New("fmpolicy FindUsingNodeID: Error iterating rows: " + err.Error())
	}
	return fmps, nil
}

func (r *fmpolicyRepository) GetInsertStr(fmpolicy fmedias.Fmpolicy) (string, error) {
	if fmpolicy.NodeID == 0 {
		return "", errors.New("fmpolicy GetInsertStr: Node ID cannot be empty or 0")
	}
	insertStr := "INSERT fmpolicy SET fmpdownload=" + strconv.Itoa(fmpolicy.FmpDownload) +
		",fmprevise=" + strconv.Itoa(fmpolicy.FmpRevise) +
		",fmpview=" + strconv.Itoa(fmpolicy.FmpView) +
		",fmpugid=" + strconv.Itoa(fmpolicy.FmpUGID) +
		",fmpugtype=" + strconv.Itoa(fmpolicy.FmpUGType) +
		",nodeid=" + strconv.Itoa(fmpolicy.NodeID)

	return insertStr, nil
}

func (r *fmpolicyRepository) GetUpdateStr(fmpolicy fmedias.Fmpolicy) (string, error) {
	if fmpolicy.NodeID == 0 || fmpolicy.FmpID == 0 {
		return "", errors.New("fmpolicy GetUpdateStr: Parameter cannot be empty")
	}
	updateStr := "UPDATE fmpolicy SET fmpdownload=" + strconv.Itoa(fmpolicy.FmpDownload) +
		",fmprevise=" + strconv.Itoa(fmpolicy.FmpRevise) +
		",fmpview=" + strconv.Itoa(fmpolicy.FmpView) +
		",fmpugid=" + strconv.Itoa(fmpolicy.FmpUGID) +
		",fmpugtype=" + strconv.Itoa(fmpolicy.FmpUGType) +
		",nodeid=" + strconv.Itoa(fmpolicy.NodeID) +
		" WHERE fmpid=" + strconv.Itoa(fmpolicy.FmpID)

	return updateStr, nil
}

func (r *fmpolicyRepository) GetDeleteStr(fmpID int) (string, error) {
	if fmpID == 0 {
		return "", errors.New("fmpolicy GetDeleteStr: Fmp ID cannot be empty or 0")
	}
	deleteStr := "DELETE FROM fmpolicy WHERE fmpid=" + strconv.Itoa(fmpID)

	return deleteStr, nil
}

func (r *convqueueRepository) Insert(convqueue fmedias.Convqueue) error {
	if convqueue.NodeID == 0 || convqueue.Convtype == "" || convqueue.FExt == "" || convqueue.FFulpath == "" || convqueue.InsDate == "" || convqueue.Priority == 0 {
		return errors.New("convqueue Insert: Parameter cannot be empty")
	}
	stmt, err := r.db.Prepare("INSERT convqueue SET nodeid=?,convtype=?,fext=?,ffulpath=?,insdate=?,priority=?")
	if err != nil {
		return errors.New("convqueue Insert: Error when preparing convqueue prepared statement: " + err.Error())
	}
	_, err = stmt.Exec(convqueue.NodeID, convqueue.Convtype, convqueue.FExt, convqueue.FFulpath, convqueue.InsDate, convqueue.Priority)
	if err != nil {
		return errors.New("convqueue Insert: Error executing convqueue statement: " + err.Error())
	}
	return nil
}

func (r *convqueueRepository) Delete(nodeID int) error {
	if nodeID == 0 {
		return errors.New("convqueue Delete: Node ID cannot be empty or 0")
	}
	stmt, err := r.db.Prepare("DELETE FROM convqueue WHERE nodeid=?")
	if err != nil {
		return errors.New("convqueue Delete: Error preparing convqueue delete statement")
	}
	_, err = stmt.Exec(nodeID)
	if err != nil {
		return errors.New("convqueue Delete: Error executing convqueue delete statement")
	}
	return nil
}

func (r *convqueueRepository) GetInsertStr(convqueue fmedias.Convqueue) (string, error) {
	if convqueue.NodeID == 0 || convqueue.Convtype == "" || convqueue.FExt == "" || convqueue.FFulpath == "" || convqueue.InsDate == "" || convqueue.Priority == 0 {
		return "", errors.New("convqueue GetInsertStr: Parameter cannot be empty")
	}

	convqueue.FFulpath = strings.Replace(convqueue.FFulpath, "\\", "\\\\", -1)

	insertStr := "INSERT convqueue SET nodeid=" + strconv.Itoa(convqueue.NodeID) +
		",convtype=\"" + convqueue.Convtype +
		"\",fext=\"" + convqueue.FExt +
		"\",ffulpath=\"" + convqueue.FFulpath +
		"\",insdate=\"" + convqueue.InsDate +
		"\",priority=" + strconv.Itoa(convqueue.Priority)

	return insertStr, nil
}

func (r *convqueueRepository) GetDeleteStr(nodeID int) (string, error) {
	if nodeID == 0 {
		return "", errors.New("convqueue GetDeleteStr: Node ID cannot be empty or 0")
	}
	deleteStr := "DELETE FROM convqueue WHERE nodeid=" + strconv.Itoa(nodeID)

	return deleteStr, nil
}

func (r *fmediaRepository) CreateFileTx(nodeStr, fmediaStr, nlStr, fverinfoStr, convStr string, fmpSlice []string) error {
	if nodeStr == "" || fmediaStr == "" || nlStr == "" || fverinfoStr == "" || convStr == "" || len(fmpSlice) == 0 {
		return errors.New("fmedia CreateFileTx: Parameter cannot be empty")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return errors.New("fmedia CreateFileTx: Error starting a transaction: " + err.Error())
	}
	nodeStmt, err := tx.Prepare(nodeStr)

	if err != nil {
		return errors.New("fmedia CreateFileTx: Error preparing node statement: " + err.Error())
	}
	fmediaStmt, err := tx.Prepare(fmediaStr)
	if err != nil {
		return errors.New("fmedia CreateFileTx: Error preparing fmedia statement: " + err.Error())
	}
	nlStmt, err := tx.Prepare(nlStr)
	if err != nil {
		return errors.New("fmedia CreateFileTx: Error preparing nodelink statement: " + err.Error())
	}
	fvStmt, err := tx.Prepare(fverinfoStr)
	if err != nil {
		return errors.New("fmedia CreateFileTx: Error preparing fverinfo statement: " + err.Error())
	}
	convStmt, err := tx.Prepare(convStr)
	if err != nil {
		return errors.New("fmedia CreateFileTx: Error preparing convqueue statement: " + err.Error())
	}
	_, err = nodeStmt.Exec()
	if err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateFileTx: Rolled-back due to node statement: " + err.Error())
	}
	_, err = fmediaStmt.Exec()
	if err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateFileTx: Rolled-back due to fmedia statement: " + err.Error())
	}
	_, err = nlStmt.Exec()
	if err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateFileTx: Rolled-back due to nodelink statement: " + err.Error())
	}
	_, err = fvStmt.Exec()
	if err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateFileTx: Rolled-back due to fverinfo statement: " + err.Error())
	}
	_, err = convStmt.Exec()
	if err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateFileTx: Rolled-back due to convqueue statement: " + err.Error())
	}
	for _, s := range fmpSlice {
		fmpStmt, e := tx.Prepare(s)
		if e != nil {
			return errors.New("fmedia CreateFileTx: Error preparing fmpolicy statement: " + err.Error())
		}
		_, err = fmpStmt.Exec()
		if err != nil {
			tx.Rollback()
			return errors.New("fmedia CreateFileTx: Rolled-back due to fmpolicy statement: " + err.Error())
		}
	}
	if err = tx.Commit(); err != nil {
		return errors.New("fmedia CreateFileTx: Error commit: " + err.Error())
	}
	return nil
}

func (r *fmediaRepository) CreateVersionTx(stmt fmedias.Statement) error {
	if stmt.CreateNodeStmt == "" || stmt.CreateFmStmt == "" || stmt.CreateNlStmt == "" ||
		stmt.CreateFvStmt == "" || stmt.CreateConvStmt == "" || len(stmt.CreateNLSlice) == 0 ||
		len(stmt.DeleteNLSlice) == 0 || len(stmt.CreateFmpSlice) == 0 {
		return errors.New("fmedia CreateVersionTx: Parameter cannot be empty")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return errors.New("fmedia CreateVersionTx: Unable to create a new transaction: " + err.Error())
	}

	createNodeS, err := tx.Prepare(stmt.CreateNodeStmt)
	if err != nil {
		return errors.New("fmedia CreateVersionTx: Unable to prepare create node statement: " + err.Error())
	}

	if _, err = createNodeS.Exec(); err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateVersionTx: Unable to execute create node statement: " + err.Error())
	}

	createFmS, err := tx.Prepare(stmt.CreateFmStmt)
	if err != nil {
		return errors.New("fmedia CreateVersionTx: Unable to prepare create fmedia statement: " + err.Error())
	}

	if _, err = createFmS.Exec(); err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateVersionTx: Unable to execute create Fmedia statement: " + err.Error())
	}

	createFvS, err := tx.Prepare(stmt.CreateFvStmt)
	if err != nil {
		return errors.New("fmedia CreateVersionTx: Unable to prepare create Fversion statement: " + err.Error())
	}

	if _, err = createFvS.Exec(); err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateVersionTx: Unable to execute create Fversion statement: " + err.Error())
	}

	for _, deleteNl := range stmt.DeleteNLSlice {
		deleteNlStmt, e := tx.Prepare(deleteNl)
		if e != nil {
			return errors.New("fmedia CreateVersionTx: Unable to prepare delete nodelink statement: " + err.Error())
		}
		if _, e = deleteNlStmt.Exec(); e != nil {
			tx.Rollback()
			return errors.New("fmedia CreateVersionTx: Unable to execute delete nodelink statemwnt: " + err.Error())
		}
	}

	for _, createNl := range stmt.CreateNLSlice {
		createNlStmt, e := tx.Prepare(createNl)
		if e != nil {
			return errors.New("fmedia CreateVersionTx: Unable to prepare create nodelink statement: " + err.Error())
		}
		if _, e = createNlStmt.Exec(); e != nil {
			tx.Rollback()
			return errors.New("fmedia CreateVersionTx: Unable to execute create nodelink statement: " + err.Error())
		}
	}

	updateCurrFmStmt, err := tx.Prepare(stmt.UpdateCurrFmStmt)
	if err != nil {
		return errors.New("fmedia CreateVersionTx: Unable to prepare update fmedia statement: " + err.Error())
	}

	if _, err = updateCurrFmStmt.Exec(); err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateVersionTx: Unable to execute update fmedia statement: " + err.Error())
	}

	updateCurrFvStmt, err := tx.Prepare(stmt.UpdateCurrFvStmt)
	if err != nil {
		return errors.New("fmedia CreateVersionTx: Unable to prepare update fversion statement: " + err.Error())
	}

	if _, err = updateCurrFvStmt.Exec(); err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateVersionTx: Unable to execute update fversion statement: " + err.Error())
	}

	for _, createFmp := range stmt.CreateFmpSlice {
		createFmpStmt, e := tx.Prepare(createFmp)
		if e != nil {
			return errors.New("fmedia CreateVersionTx: Unable to prepare create fmpolicy statement: " + err.Error())
		}
		if _, e = createFmpStmt.Exec(); e != nil {
			tx.Rollback()
			return errors.New("fmedia CreateVersionTx: Unable to execute create fmpoliy statement: " + err.Error())
		}
	}

	createConvStmt, err := tx.Prepare(stmt.CreateConvStmt)
	if err != nil {
		return errors.New("fmedia CreateVersionTx: Unable to prepare create convqueue statement: " + err.Error())
	}
	if _, err = createConvStmt.Exec(); err != nil {
		tx.Rollback()
		return errors.New("fmedia CreateVersionTx: Unable to execute create convqueue statement: " + err.Error())
	}

	if err = tx.Commit(); err != nil {
		return errors.New("fmedia CreateVersionTx: Unable to commit" + err.Error())
	}

	return nil
}
