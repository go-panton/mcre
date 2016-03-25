package users

import (
	"time"

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
	createdAt    time.Time
	userID       int
	sessionToken string
}

func makeSignUpEndPoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		err := svc.User(req.Username, req.Password)
		if err != nil {
			return SignUpResponse{time.Now(), 0, ""}, nil
		}
		return SignUpResponse{time.Now(), 1, "xDXDXDXDXD"}, nil
	}
}
