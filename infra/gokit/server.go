package gokit

import "github.com/gorilla/mux"

// Server is the interface that serves
type Server interface {
	// RouteTo connects internal server-routes to an external subrouter, returns
	// router to chain actions.
	//
	// if subRouter is nil, returns error or you might want to initialize a
	// default mux router instead.
	RouteTo(router *mux.Router) *mux.Router
}
