package mcre

import (
	"fmt"
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

// TODO: come back to put object into database instead of fs.
// TODO: might want to allow user to pre-set target directory.

// PutObject mkdirs a root directory with bucket, creates a file with
// bucket/key and writes file content with body from http-request, and returns
// the etag from hashing the file.
func (mcre *MCRE) PutObject(bucket, key string, content io.Reader) (string, error) {

	// TODO: Existing file issue.
	if err := os.MkdirAll(bucket, os.ModePerm); err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("create root-dir %v: %v", bucket, err)
	}

	file, err := os.OpenFile(filepath.Join(bucket, key), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("create file %v: %v", key, err)
	}

	if _, err := io.Copy(file, content); err != nil {
		return "", fmt.Errorf("create file %v: %v", key, err)
	}

	return "etag-to-be-implemented", nil
}
