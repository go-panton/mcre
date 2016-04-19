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

func ConnectDatabase(dbName, colName string) *mgo.Collection {

	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
	}
	//DB for database name C for collections which equivalent to tables in relational database
	return session.DB(dbName).C(colName)
}

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
