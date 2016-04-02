package userserver

import (
	"testing"

	"github.com/go-panton/mcre/infra/store/mongo"
	"github.com/go-panton/mcre/users"
	"github.com/go-panton/mcre/users/model"
	"github.com/gorilla/mux"

	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"golang.org/x/net/context"
)

var tests = []models.User{
	{"alex", ""},     //password empty
	{"", "root"},     //username empty
	{"", ""},         //both field empty
	{"alex", "root"}, //username already exist in database
	{"Rex", "Gear"},  //success case
}

func TestServer(t *testing.T) {
	repo := mongo.NewMockUserRepository()

	for _, pair := range tests {
		bytePair, err := json.Marshal(pair)
		if err != nil {
			fmt.Println("Error to marshal json")
		}
		w := httptest.NewRecorder()

		r, err := http.NewRequest("POST", "/users", bytes.NewReader(bytePair))
		r.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Fatal(err)
		}

		server(repo).ServeHTTP(w, r)

		fmt.Println(w.Code)
		fmt.Println(w.Body.String())
	}

}

func server(repo models.UserRepository) http.Handler {
	server := NewServer(context.Background(), users.NewService(repo))
	router := mux.NewRouter()
	return server.RouteTo(router)
}
