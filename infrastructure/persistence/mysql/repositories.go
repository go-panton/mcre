package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-panton/mcre/users/model"
	"database/sql"
)

type userRepository struct{
	db *sql.DB
}

func ConnectDatabase(connString string) *sql.DB{
	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Println("Error connecting to database")
	}
	return db
}

func NewUser(db *sql.DB) models.UserRepository{
	return &userRepository{db}
}

func (r *userRepository)Insert(username,password string) error{
	insStat, err := r.db.Prepare("INSERT user SET username=?,password=?")

	if err != nil {
		return err
	}
	_, err1 := insStat.Exec(username,password)

	if err1 != nil {
		return err1
	}
	return nil
}

func (r *userRepository)Find(username string) (*models.User, error) {
	var resultName,password string
	err := r.db.QueryRow("SELECT * FROM user WHERE username=?",username).Scan(&resultName,&password)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	default:
		return &models.User{resultName,password},nil
	}
}