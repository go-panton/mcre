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

func ConnectDatabase(connString string){
	db, err := sql.Open("mysql", connString)
	defer db.Close()

	if err != nil {
		fmt.Println("Error connecting to database")
	}
}

func NewUser(db *sql.DB) models.UserRepository{
	return &userRepository{db}
}

func (r *userRepository)Insert(username,password string) error{
	userid := 35325//TODO: remove userid field in mysql database

	insStat, err := r.db.Prepare("INSERT user SET username=?,password=?,userid=?")

	if err != nil {
		return err
	}
	_, err1 := insStat.Exec(username,password,userid)

	if err1 != nil {
		return err1
	}
	return nil
}

func (r *userRepository)Find(userID string) (*models.User, error) {
	username := "alex"
	password := "root"
	err := r.db.QueryRow("SELECT * FROM user WHERE username=? AND password=?",username,password).Scan(&models.User{})
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	default:
		return &models.User{},nil
	}
}