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
	CreatedAt    time.Time `json:"createdat"`
	UserID       int       `json:"userid"`
	SessionToken string    `json:"sessiontoken"`
}

func makeSignUpEndPoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		err := svc.SignUp(req.Username, req.Password)
		if err != nil {
			return SignUpResponse{time.Now(), 0, ""}, nil
		}
		return SignUpResponse{time.Now(), 1, "xDXDXDXDXD"}, nil
	}
}
