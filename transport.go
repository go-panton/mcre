package mcre

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// PutObjectInput is the input of PutObject
type PutObjectInput struct {
	Bucket string
	Key    string
	Body   io.ReadCloser
}

// PutObjectOutput is the output of PutObject
type PutObjectOutput struct {
	//TODO: check if etag is important
	Etag      string `json:"-"` //unexported as this will not be encoded in body
	VersionID string `json:"version_id,omitempty"`
}

// decodePutObject decodes path from request, assigns values into PutObjectInput
// struct.
//
// Note:
//  - Body shall not be parsed as JSON. It is file content.
//  - Path will be dissected as bucket, key.
func decodePutObject(request *http.Request) (interface{}, error) {
	vars := mux.Vars(request)
	bucket := vars["bucket"]
	key := vars["key"]

	switch {
	case bucket == "":
		return nil, fmt.Errorf("empty bucket")
	case key == "":
		return nil, fmt.Errorf("empty key")
	}

	return &PutObjectInput{Bucket: bucket, Key: key, Body: request.Body}, nil
}

// TODO: add tests for Etag generation.
// makePutObjectEndpoint creates an Endpoint from mcre service, the endpoint
// access to background context, asserts input as PutObjectInput, calls
// PutObject API and returns PutObjectOutput.
func makePutObjectEndpoint(mcre *MCRE) endpoint.Endpoint {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		in := input.(PutObjectInput)
		etag, err := mcre.PutObject(in.Bucket, in.Key, in.Body)
		if err != nil {
			return nil, err // this returns error http response, how do u want it? TODO: study this.
		}
		return &PutObjectOutput{Etag: etag}, nil
	}
}

// encodePutObject packs info to be sent over netork.
//
// TODO: content length?
//
// Note:
//	- no body content.
//	- set header["etag"]
//  -
func encodePutObject(w http.ResponseWriter, output interface{}) error {
	out := output.(PutObjectOutput)
	w.Header().Set("ETag", out.Etag)

	return nil
}
