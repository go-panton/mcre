package mcre

import (
	"io"
	"os"
	"path/filepath"
)

// MCRE is Multi-Content Repository that provides simple file storage.
// TODO: make an api-interface
type MCRE struct{}

// PutObjectInput is the input of PutObject
type PutObjectInput struct {
	Bucket string
	Key    string
	Body   io.Reader
}

// PutObjectOutput is the output of PutObject
type PutObjectOutput struct {
	//TODO: check if etag is important
	Etag      string `json:"-"` //unexported as this will not be encoded in body
	VersionID string `json:"version_id,omitempty"`
}

// PutObject adds an object to a bucket
// TODO: come back to put object into database instead of fs.
// TODO: might want to allow user to pre-set target directory.
// TODO: to have custom error message
func (mcre *MCRE) PutObject(input *PutObjectInput) (output *PutObjectOutput, err error) {

	dir := input.Bucket
	file := input.Key
	content := input.Body

	// TODO: Existing file issue.
	if err := os.MkdirAll(input.Bucket, os.ModePerm); err != nil && !os.IsExist(err) {
		return nil, err
	}

	fd, err := os.OpenFile(filepath.Join(dir, file), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(fd, content); err != nil {
		return nil, err
	}

	output.Etag = "ETAG-to-be-implemented"

	return
}
