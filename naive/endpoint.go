package naive

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

// MakeEndpoint creates an Endpoint from mcre service, the endpoint
// access to background context, asserts input as PutObjectInput, calls
// PutObject API and returns PutObjectOutput.
// TODO: add tests for Etag generation.
func MakeEndpoint(mcre *naive) endpoint.Endpoint {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		in := input.(PutObjectInput)
		etag, err := mcre.PutObject(in.Bucket, in.Key, in.Body)
		if err != nil {
			return nil, err // this returns error http response, how do u want it? TODO: study this.
		}
		return &PutObjectOutput{Etag: etag}, nil
	}
}
