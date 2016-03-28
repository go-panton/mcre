package files

import (
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type badRequestError error

// MakeHandler returns a handler for the micre service.
func MakeHandler(ctx context.Context, svc Service) http.Handler {

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	simpleDownloadHandler := kithttp.NewServer(
		ctx,
		makeDownloadEndpoint(svc),
		decodeDownloadRequest,
		encodeDownloadResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/files/{fileid}", simpleDownloadHandler).Methods("GET")

	return r
}

// encodeError is an adapter function that writes HTTP error response, with
// message from err, and status-code according to error type.
func encodeError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case badRequestError:
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
