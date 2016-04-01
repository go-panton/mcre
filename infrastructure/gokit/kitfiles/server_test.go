package kitfiles

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-panton/mcre/files"

	"golang.org/x/net/context"
)

func TestIntegration(t *testing.T) {
	server := MakeHandler(context.Background(), files.NewService())

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/files/123", nil)
	if err != nil {
		log.Fatal(err)
	}
	server.ServeHTTP(w, r)

	fmt.Println(w.Body.String())
}
