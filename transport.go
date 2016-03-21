package mcre

import (
	"fmt"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

// MakeHandler creates gokit-handler per API with context, mcre-service and
// logger, then mux maps url to each respective api-handler, finally mux returns
// as an api-router.
func MakeHandler(ctx context.Context, s Service, logger kitlog.Logger) http.Handler {

	putObjectHandler := kithttp.NewServer(
		ctx,
		MakePutObjectEndpoint(s),
		DecodePutObject,
		EncodePutObject,
	)
	r := mux.NewRouter()
	r.Handle("/v1/{bucket}/{key}", putObjectHandler)
	return r
}

// DecodePutObject decodes path from request, assigns values into PutObjectInput
// struct.
//
// Note:
//  - Body shall not be parsed as JSON. It is file content.
//  - Path will be dissected as bucket, key.
func DecodePutObject(request *http.Request) (interface{}, error) {
	vars := mux.Vars(request)
	bucket := vars["bucket"]
	key := vars["key"]

	switch {
	case bucket == "":
		return nil, fmt.Errorf("empty bucket")
	case key == "":
		return nil, fmt.Errorf("empty key")
	}

	defer request.Body.Close()

	return &PutObjectInput{Bucket: bucket, Key: key, Body: request.Body}, nil
}

// EncodePutObject packs info to be sent over netork.
//
// TODO: content length?s
//
// Note:
//	- no body content.
//	- set header["etag"]
//  -
func EncodePutObject(w http.ResponseWriter, output interface{}) error {
	out := output.(PutObjectOutput)
	w.Header().Set("ETag", out.Etag)

	return nil
}
