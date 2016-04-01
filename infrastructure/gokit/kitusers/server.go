package kitusers

import (
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-panton/mcre/users"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

//MakeHandler create a func to handle that url
func MakeHandler(ctx context.Context, svc users.Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	signUpHandler := kithttp.NewServer(
		ctx,
		makeSignUpEndPoint(svc),
		decodeSignUpRequest,
		encodeSignUpResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/users", signUpHandler).Methods("POST")

	return r
}
