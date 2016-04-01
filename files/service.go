package files

import (
	"io"
	"strings"
)

type BadRequestError error

// Service is the interface of download API
type Service interface {
	//Download(simple) opens and reads file based on fileid
	Download(fileid string) io.Reader
}

// NewService instantiates new download-service.
func NewService() Service {
	return &service{}
}

type service struct{}

func (svc *service) Download(fileid string) io.Reader {
	return strings.NewReader("This is your requested file content.")
}
