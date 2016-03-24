package main

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-panton/mcre/download/v1"
	"github.com/go-panton/mcre/user/v1"
)

func main() {

	ds := download.NewService()
	us := user.NewService()
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("/download/v1/", download.MakeHandler(ctx, ds))
	mux.Handle("/user/v1", user.MakeHandler(ctx, us))

	http.Handle("/", mux)
	log.Fatal(http.ListenAndServe(":8282", nil))
}

// type server struct{}
//
// func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "woi")
// }
