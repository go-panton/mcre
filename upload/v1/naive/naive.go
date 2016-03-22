package naive

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-panton/mcre"
)

type naive struct {
}

// New instantiates new serviceAPI object
func New() mcre.Service {
	return &naive{}
}

// PutObject mkdirs a root directory with bucket, creates a file with
// bucket/key and writes file content with body from http-request, and returns
// the etag from hashing the file.
//
// TODO: come back to put object into database instead of fs.
// TODO: might want to allow user to pre-set target directory.
func (s *naive) PutObject(bucket, key string, content io.Reader) (string, error) {

	// TODO: Existing file issue.
	if err := os.MkdirAll(bucket, os.ModePerm); err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("create root-dir %v: %v", bucket, err)
	}

	if filepath.IsAbs(key) {
		return "", fmt.Errorf("path is absolute: %v", key)
	}

	keyPath := filepath.Dir(key)
	if keyPath != "." {
		if err := os.MkdirAll(filepath.Join(bucket, keyPath), os.ModePerm); err != nil && !os.IsExist(err) {
			return "", fmt.Errorf("create dir for key %v: %v", key, err)
		}
	}

	file, err := os.OpenFile(filepath.Join(bucket, key), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("create file %v: %v", key, err)
	}

	if _, err := io.Copy(file, content); err != nil {
		return "", fmt.Errorf("write file %v: %v", key, err)
	}

	return "etag-to-be-implemented", nil
}
