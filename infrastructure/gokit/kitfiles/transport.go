package kitfiles

import (
	"io"
	"net/http"

	"github.com/go-panton/mcre/files"
	"github.com/gorilla/mux"
)

// /files/:fileid
func decodeDownloadRequest(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fid := vars["fileid"]

	return files.DownloadRequest{FileId: fid}, nil
}

func encodeDownloadResponse(w http.ResponseWriter, response interface{}) error {
	resp := response.(files.DownloadResponse)
	w.Header().Set("etag", "whatisetag")
	io.Copy(w, resp.FileContent)
	return nil
}
