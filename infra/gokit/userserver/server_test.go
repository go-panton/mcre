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

	"errors"

	"golang.org/x/net/context"
)

type testPair struct {
	TestData models.User
	Expected error
}

var tests = []testPair{
	{models.User{"alex", ""}, users.BadRequestError(errors.New("The password is empty."))},
	{models.User{"", "root"}, users.BadRequestError(errors.New("The username is empty."))},
	{models.User{"", ""}, users.BadRequestError(errors.New("The username is empty."))},
	{models.User{"alex", "root"}, users.BadRequestError(errors.New("The username has already been taken."))},
	{models.User{"Rex", "Gear"}, nil},
}

func TestServer(t *testing.T) {
	repo := mongo.NewMockUserRepository()

	for _, test := range tests {
		bytePair, err := json.Marshal(test)
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
