package user

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

//SignUpRequest is a request struct
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//SignUpResponse is a response struct
type SignUpResponse struct {
	Status bool
	Err    string
}

func makeSignUpEndPoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		check, err := svc.User(req.Username, req.Password)
		if err != nil {
			return SignUpResponse{check, err.Error()}, nil
		}
		return SignUpResponse{check, ""}, nil
	}
}
