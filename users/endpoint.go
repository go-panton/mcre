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
	CreatedAt    time.Time `json:"createdAt"`
	//UserID       int       `json:"userid"`
	//SessionToken string    `json:"sessiontoken"`
}

type ErrorResponse struct {
	Code        int     `json:"code"`
	Message     string  `json:"message"`
}

func makeSignUpEndPoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		err := svc.SignUp(req.Username, req.Password)
		if err != nil {
			return SignUpResponse{time.Now()}, err
		}
		return SignUpResponse{time.Now()}, nil
	}
}
