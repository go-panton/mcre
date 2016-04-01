package users

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	kithttp "github.com/go-kit/kit/transport/http"
)

//MakeHandler create a func to handle that url
func MakeHandler(ctx context.Context, svc Service) http.Handler {
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

func decodeSignUpRequest(r *http.Request) (interface{}, error) {
	var req SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return SignUpRequest{req.Username, req.Password}, nil
}

func encodeSignUpResponse(w http.ResponseWriter, response interface{}) error {
	resp := response.(SignUpResponse)
	respb, err := json.Marshal(resp)
	if err != nil {
		return errors.New("Unable to marshal repsonse into json")
	}
	w.Header().Add("Status", http.StatusText(http.StatusCreated))
	w.Header().Add("Location", "http://localhost:8282/users/v1/1")
	w.Write(respb)
	return nil
}

func encodeError(w http.ResponseWriter, err error){
	switch err.(type) {
	case kithttp.BadRequestError:
		http.Error(w,err.Error(), http.StatusBadRequest)
	default:
		http.Error(w,err.Error(), http.StatusInternalServerError)
	}
}