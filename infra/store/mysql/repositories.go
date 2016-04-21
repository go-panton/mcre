package mysql

import (
	"fmt"
	"strconv"

	"database/sql"

	"errors"

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

	stmt, err := r.db.Prepare("DELETE FROM node where nodeid=?")
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
	if node.NodeID == 0 || node.NodeDesc == "" || node.NodeGID == 0 {
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
