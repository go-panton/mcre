package mysql

import (
	"fmt"

	"database/sql"

	"errors"

	"github.com/go-panton/mcre/users/model"
	_ "github.com/go-sql-driver/mysql"
)

type userRepository struct {
	db *sql.DB
}

type seqRepository struct {
	db *sql.DB
}

func ConnectDatabase(connString string) *sql.DB {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Println("Error connecting to database")
	}
	return db
}

func NewUser(db *sql.DB) models.UserRepository {
	return &userRepository{db}
}

func NewSeq(db *sql.DB) models.SeqRepository {
	return &seqRepository{db}
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

func (r *userRepository) Find(username string) (*models.User, error) {
	var resultName, resultPassword string
	err := r.db.QueryRow("SELECT * FROM user WHERE username=?", username).Scan(&resultName, &resultPassword)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	default:
		return &models.User{resultName, resultPassword}, nil
	}
}

func (r *userRepository) Verify(username, password string) (*models.User, error) {
	var resultName, resultPassword string
	err := r.db.QueryRow("SELECT * FROM user WHERE username=? AND password=?").Scan(&resultName, &resultPassword)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	default:
		return &models.User{resultName, password}, nil
	}
}

//Get get the value from seqtbl
func (r *seqRepository) Get(query string) (int, error) {
	if query == "" {
		return 0, errors.New("The query string is empty.")
	}

	var value int
	err := r.db.QueryRow("SELECT pseq FROM seqtbl WHERE PNAME=?", query).Scan(&value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

//Update update the value of table seqtbl
func (r *seqRepository) Update(value int, query string) error {
	if value < 1 {
		return errors.New("The update value should never be less than 1")
	}
	if query == "" {
		return errors.New("The query string is empty.")
	}

	stmt, err := r.db.Prepare("UPDATE seqtbl SET pseq=? WHERE pname=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(value, query)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("Rows affected: ", affect)
	return nil
}
