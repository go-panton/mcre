package mcre

import "io"

// PutObjectInput is the payload of PutObject
type PutObjectInput struct {
	Bucket string
	Key    string
	Body   io.Reader
}

// PutObjectOutput is the payload of PutObject
type PutObjectOutput struct {
	//TODO: check if etag is important
	Etag      string `json:"-"` //unexported as this will not be encoded in body
	VersionID string `json:"version_id,omitempty"`
}
