package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-panton/mcre/files"
	"github.com/go-panton/mcre/infra/gokit/fileserver"
	"github.com/go-panton/mcre/infra/gokit/userserver"
	"github.com/go-panton/mcre/infra/store/mysql"
	"github.com/go-panton/mcre/users"
	"github.com/gorilla/mux"
)

var (
	port   = flag.String("port", ":8282", "Listen port")
	router = mux.NewRouter()
	mcrev1 = router.PathPrefix("/mcre/v1").Subrouter()
)

func main() {
	flag.Parse()

	//mongoDbName := "go_panton"
	//mongoColName := "user"
	mysqlconnectionString := "root:root123@/go_panton"

	//us := users.NewService(mongo.NewUser(mongo.ConnectDatabase(mongoDbName, mongoColName)))

	us := users.NewService(mysql.NewUser(mysql.ConnectDatabase(mysqlconnectionString)))
	usersvr := userserver.NewServer(context.Background(), us)
	usersvr.RouteTo(mcrev1)

	filesvr := fileserver.NewServer(context.Background(), files.NewService())
	filesvr.RouteTo(mcrev1)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(*port, nil))
}
