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
	signUpHandler := kithttp.NewServer(
		ctx,
		makeSignUpEndPoint(svc),
		decodeSignUpRequest,
		encodeSignUpResponse,
	)

	r := mux.NewRouter()
	r.Handle("/users/v1", signUpHandler).Methods("POST")

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
	w.Header().Add("Location", "http://localhost:8282/users/XDXDXD")
	w.Write(respb)
	return nil
}
