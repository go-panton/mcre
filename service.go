package mcre

import "io"

// Service is an interface of MCRE
type Service interface {
	PutObject(bucket, key string, content io.Reader) (string, error)
}
