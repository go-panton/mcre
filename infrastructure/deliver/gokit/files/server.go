package files

// MakeHandler returns a handler for the micre service.
import (
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-panton/mcre/files"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func MakeHandler(ctx context.Context, svc files.Service) http.Handler {

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	simpleDownloadHandler := kithttp.NewServer(
		ctx,
		makeDownloadEndpoint(svc),
		decodeDownloadRequest,
		encodeDownloadResponse,
		opts...,
	)

	r := mux.NewRouter()
	s := r.PathPrefix("/files").Subrouter()
	s.Handle("/{fileid}", simpleDownloadHandler).Methods("GET")

	return r
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
