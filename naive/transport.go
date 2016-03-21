package naive

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

	return &PutObjectInput{Bucket: bucket, Key: key, Body: request.Body}, nil
}
