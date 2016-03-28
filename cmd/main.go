package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-panton/mcre/download/v1"
	"github.com/go-panton/mcre/users/v1"
)

var (
	httpAddr = flag.String("http", ":8282", "Listen address")
)

func main() {
	flag.Parse()

	ds := download.NewService()
	us := users.NewService()
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("/download/v1/", download.MakeHandler(ctx, ds))
	mux.Handle("/users/v1", users.MakeHandler(ctx, us))

	http.Handle("/", mux)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

// type server struct{}
//
// func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "woi")
// }
