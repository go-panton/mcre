// Package download contains the following:
//
// encode
// decode
// makehandler
package download

import (
	"io"
	"net/http"

	"golang.org/x/net/context"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
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
	r.Handle("/download/v1/{fileid}", simpleDownloadHandler).Methods("GET")

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

// @/download/v1/:fileid
func decodeDownloadRequest(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fid := vars["fileid"]

	return downloadRequest{fileid: fid}, nil
}

func encodeDownloadResponse(w http.ResponseWriter, response interface{}) error {
	resp := response.(downloadResponse)
	w.Header().Set("etag", "whatisetag")
	io.Copy(w, resp.filecontent)
	return nil
}
