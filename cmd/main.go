package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-panton/mcre/infrastructure/gokit"
)

var (
	port = flag.String("port", ":8282", "Listen port")
)

func main() {
	flag.Parse()

	http.Handle("/", gokit.NewKit())
	log.Fatal(http.ListenAndServe(*port, nil))
}
