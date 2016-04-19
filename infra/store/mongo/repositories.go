package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"

	"github.com/go-panton/mcre/users/model"
)

type userRepository struct {
	col *mgo.Collection
}

//ConnectDatabase returns connection of mongo database based on dbName and colName provided
func ConnectDatabase(dbName, colName string) *mgo.Collection {

	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
	}
	//DB for database name C for collections which equivalent to tables in relational database
	return session.DB(dbName).C(colName)
}

//NewUser return a userRepository based on mongo Collection provided
func NewUser(col *mgo.Collection) model.UserRepository {
	return &userRepository{col}
}

func (r *userRepository) Insert(username, password string) error {
	newUser := model.User{Username: username, Password: password}

	err := r.col.Insert(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Find(username string) (*model.User, error) {
	result := &model.User{}
	err := r.col.Find(bson.M{"username": username}).One(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepository) Verify(username, password string) (*model.User, error) {
	result := &model.User{}
	err := r.col.Find(bson.M{"username": username, "password": password}).One(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
