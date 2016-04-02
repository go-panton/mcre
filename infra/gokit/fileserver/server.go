package fileserver

// MakeHandler returns a handler for the micre service.
import (
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/go-panton/mcre/files"
	"github.com/go-panton/mcre/infra/gokit"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type fileServer struct {
	ctx context.Context
	svc files.Service
}

// NewServer returns instance that hosts file services.
func NewServer(ctx context.Context, svc files.Service) gokit.Server {
	return &fileServer{ctx: ctx, svc: svc}
}

func (fs *fileServer) RouteTo(router *mux.Router) *mux.Router {

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	simpleDownloadHandler := kithttp.NewServer(
		fs.ctx,
		makeDownloadEndpoint(fs.svc),
		decodeDownloadRequest,
		encodeDownloadResponse,
		opts...,
	)
	router.Handle("/files/{fileid}", simpleDownloadHandler).Methods("GET")

	return router
}

// encodeError is an adapter function that writes HTTP error response, with
// message from err, and status-code according to error type.
func encodeError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case files.BadRequestError:
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
