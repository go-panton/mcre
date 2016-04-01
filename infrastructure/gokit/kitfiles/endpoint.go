package kitfiles

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-panton/mcre/files"
	"golang.org/x/net/context"
)

func makeDownloadEndpoint(svc files.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(files.DownloadRequest)
		resp := svc.Download(req.FileId)
		return files.DownloadResponse{FileContent: resp}, nil
	}
}
