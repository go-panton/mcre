package kitusers

import (
	"testing"
	"github.com/go-panton/mcre/users"
	"github.com/go-panton/mcre/users/model"

	"golang.org/x/net/context"
	"github.com/go-panton/mcre/infrastructure/persistence/mongo"
	"net/http/httptest"
	"net/http"
	"log"
	"fmt"
	"bytes"
	"encoding/json"
)

var tests = []models.User{
	{"alex", ""},//password empty
	{"","root"},//username empty
	{"",""},//both field empty
	{"alex","root"},//username already exist in database
	{"Rex","Gear"},//success case
}

func TestServer (t *testing.T){
	repo := mongo.NewMockUserRepository()

	server := MakeHandler(context.Background(),users.NewService(repo))

	for _, pair := range tests {
		bytePair, err := json.Marshal(pair)
		if err != nil {
			fmt.Println("Error to marshal json")
		}
		w := httptest.NewRecorder()

		r, err := http.NewRequest("POST", "/users", bytes.NewReader(bytePair))
		r.Header.Set("Content-Type","application/json")
		if err != nil {
			log.Fatal(err)
		}

		server.ServeHTTP(w,r)

		fmt.Println(w.Code)
		fmt.Println(w.Body.String())
	}

}


