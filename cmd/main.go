package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-panton/mcre/files"
	"github.com/go-panton/mcre/users"
)

var (
	port = flag.String("port", ":8282", "Listen port")
)

func main() {
	flag.Parse()

	fs := files.NewService()
	us := users.NewService()
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("/mcre/v1/files/", files.MakeHandler(ctx, fs))
	mux.Handle("/mcre/v1/users/", users.MakeHandler(ctx, us))

	http.Handle("/", mux)
	log.Fatal(http.ListenAndServe(*port, nil))
}
