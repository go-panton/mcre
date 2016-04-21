package mongo

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestMongo(t *testing.T) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Couldn't connect to mongo database" + err.Error())
	}

	//DB for database name C for collections which equivalent to tables in relational database
	user := session.DB("go_panton").C("user")

	err = NewUser(user).Insert("alex", "213")
	if err != nil {
		t.Errorf("Error inserting data into collection" + err.Error())
	}
}
