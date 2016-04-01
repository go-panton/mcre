package gokit

import (
	"net/http"

	"github.com/go-panton/mcre/files"
	"github.com/go-panton/mcre/infrastructure/gokit/kitfiles"
	"github.com/go-panton/mcre/infrastructure/gokit/kitusers"
	"github.com/go-panton/mcre/infrastructure/persistence/mysql"
	"github.com/go-panton/mcre/users"
	"golang.org/x/net/context"
)

// NewKit returns kit handlers
func NewKit() http.Handler {
	//mongoDbName := "go_panton"
	//mongoColName := "user"
	mysqlconnectionString := "root:root123@/go_panton"

	fs := files.NewService()

	//us := users.NewService(mongo.NewUser(mongo.ConnectDatabase(mongoDbName,mongoColName)))
	us := users.NewService(mysql.NewUser(mysql.ConnectDatabase(mysqlconnectionString)))

	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("/mcre/v1/files/", kitfiles.MakeHandler(ctx, fs))
	mux.Handle("/mcre/v1/users/", kitusers.MakeHandler(ctx, us))

	return mux
}
