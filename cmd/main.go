package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-panton/mcre/files"
	"github.com/go-panton/mcre/users"

	"github.com/go-panton/mcre/infrastructure/persistence/mongo"

	"github.com/go-panton/mcre/infrastructure/persistence/mysql"
)

var (
	port = flag.String("port", ":8282", "Listen port")
)

func main() {
	flag.Parse()

	//mongoDbName := "go_panton"
	//mongoColName := "user"
	mysqlconnectionString := "root:root123@/go_panton"

	fs := files.NewService()

	//us := users.NewService(mongo.NewUser(mongo.ConnectDatabase(mongoDbName,mongoColName)))
	us := users.NewService(mysql.NewUser(mysql.ConnectDatabase(mysqlconnectionString)))

	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("/mcre/v1/files/", files.MakeHandler(ctx, fs))
	mux.Handle("/mcre/v1/users/", users.MakeHandler(ctx, us))

	http.Handle("/", mux)
	log.Fatal(http.ListenAndServe(*port, nil))
}
