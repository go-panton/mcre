package fileserver

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-panton/mcre/files"
	"github.com/gorilla/mux"

	"golang.org/x/net/context"
)

func TestIntegration(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/files/123", nil)
	if err != nil {
		log.Fatal(err)
	}
	server().ServeHTTP(w, r)

	fmt.Println(w.Body.String())
}

func server() http.Handler {
	server := NewServer(context.Background(), files.NewService())
	router := mux.NewRouter()
	return server.RouteTo(router)
}
