// Package download contains the following:
//
// encode
// decode
// makehandler
package download

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHandler returns a handler for the micre service.
func MakeHandler(ctx context.Context, svc Service) http.Handler {
	simpleDownloadHandler := kithttp.NewServer(
		ctx,
		makeDownloadEndpoint(svc),
		decodeDownloadRequest,
		encodeDownloadResponse,
	)

	r := mux.NewRouter()
	r.Handle("/download/v1/{fileid}", simpleDownloadHandler).Methods("GET")
	//r.Handle("/", simpleDownloadHandler).Methods("GET")

	fmt.Println("MakeHandler")

	return r
}

// @/download/v1/:fileid
func decodeDownloadRequest(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fid, ok := vars["fileid"]

	if !ok {
		return nil, fmt.Errorf("error: %v", "bad route")
	}

	return downloadRequest{fileid: fid}, nil
}

func encodeDownloadResponse(w http.ResponseWriter, response interface{}) error {
	resp := response.(downloadResponse)
	io.Copy(w, resp.filecontent)
	return nil
}
