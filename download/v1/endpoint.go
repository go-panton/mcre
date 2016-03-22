package download

import (
	"io"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type downloadRequest struct {
	fileid string
}

type downloadResponse struct {
	filecontent io.Reader
}

func makeDownloadEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(downloadRequest)
		resp := svc.Download(req.fileid)
		return downloadResponse{filecontent: resp}, nil
	}
}
