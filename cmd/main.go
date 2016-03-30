package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-panton/mcre/files"
	"github.com/go-panton/mcre/users"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/go-panton/infrastructure/persistence/mongo"
	"github.com/go-panton/infrastructure/persistence/mysql"
)

var (
	port = flag.String("port", ":8282", "Listen port")
)

func main() {
	flag.Parse()
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
	}

	//DB for database name C for collections which equivalent to tables in relational database
	mongoDbName := "go_panton"
	mongoColName := "users"
	mysqlconnectionString := "root:root123@/go_panton"

	fs := files.NewService()
	//temporary check to switch between different database
	if mongoDbName != ""{
		us := users.NewService(mongo.NewUser(mongo.ConnectDatabase(mongoDbName,mongoColName)))
	}
	us := users.NewService(mysql.NewUser(mysql.ConnectDatabase(mysqlconnectionString)))
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("/mcre/v1/files/", files.MakeHandler(ctx, fs))
	mux.Handle("/mcre/v1/users/", users.MakeHandler(ctx, us))

	http.Handle("/", mux)
	log.Fatal(http.ListenAndServe(*port, nil))
}
