package userserver

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-panton/mcre/infra/gokit"
	"github.com/go-panton/mcre/users"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type userServer struct {
	ctx context.Context
	svc users.Service
}

// NewServer returns instance that hosts file services.
func NewServer(ctx context.Context, svc users.Service) gokit.Server {
	return &userServer{ctx: ctx, svc: svc}
}

func (svr *userServer) RouteTo(router *mux.Router) *mux.Router {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	signUpHandler := kithttp.NewServer(
		svr.ctx,
		makeSignUpEndPoint(svr.svc),
		decodeSignUpRequest,
		encodeSignUpResponse,
		opts...,
	)
	router.Handle("/users", signUpHandler).Methods("POST")
	return router
}
