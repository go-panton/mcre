package userserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-panton/mcre/users"
)

func decodeSignUpRequest(r *http.Request) (interface{}, error) {
	var req users.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func encodeSignUpResponse(w http.ResponseWriter, response interface{}) error {
	resp := response.(users.SignUpResponse)
	respb, err := json.Marshal(resp)
	if err != nil {
		return errors.New("Unable to marshal response into json")
	}
	w.Header().Add("Status", http.StatusText(http.StatusCreated))
	w.Header().Add("Location", "http://localhost:8282/users/v1/1")
	w.Write(respb)
	return nil
}

func encodeError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case users.BadRequestError:
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
