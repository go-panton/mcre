package userserver

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-panton/mcre/users"
	"golang.org/x/net/context"
)

func makeSignUpEndPoint(svc users.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(users.SignUpRequest)
		err := svc.SignUp(req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return users.SignUpResponse{CreatedAt: time.Now()}, nil
	}
}
