package mongo

import (
	"fmt"
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestMongo(t *testing.T) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
	}

	//DB for database name C for collections which equivalent to tables in relational database
	user := session.DB("go_panton").C("user")

	NewUser(user).Insert("alex", "213")

}
