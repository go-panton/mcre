package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"

	"github.com/go-panton/mcre/users/models"
)

type userRepository struct {
	col *mgo.Collection
}

//ConnectDatabase return a collection based on dbname and colname provided
func ConnectDatabase(dbName, colName string) *mgo.Collection {

	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
	}
	//DB for database name C for collections which equivalent to tables in relational database
	return session.DB(dbName).C(colName)
}

//NewUser return a userRepository based on mongo collection provided
func NewUser(col *mgo.Collection) models.UserRepository {
	return &userRepository{col}
}

func (r *userRepository) Insert(username, password string) error {
	newUser := models.User{Username: username, Password: password}

	err := r.col.Insert(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Find(username string) (*models.User, error) {
	result := &models.User{}
	err := r.col.Find(bson.M{"username": username}).One(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
