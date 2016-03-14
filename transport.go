package mcre

import (
	"io"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
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

// decodePutObject decodes http-request into PutObjectInput struct.
func decodePutObject(r *http.Request) (request interface{}, err error) {
	//mux.Vars()
	return nil, nil
}

// TODO: add tests for Etag generation.
// makePutObjectEndpoint creates an Endpoint from mcre service, the endpoint
// access to background context, asserts request as PutObjectRequest, calls
// PutObject API and returns PutObjectOutput.
func makePutObjectEndpoint(mcre *MCRE) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		in := request.(PutObjectInput)
		etag, err := mcre.PutObject(in.Bucket, in.Key, in.Body)
		if err != nil {
			return nil, err // this returns error http response. TODO: study this.
		}
		return &PutObjectOutput{Etag: etag}, nil
	}
}
func encodePutObject(w http.ResponseWriter, response interface{}) error {
	return nil
}
