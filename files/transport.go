package files

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// @/download/v1/:fileid
func decodeDownloadRequest(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fid := vars["fileid"]

	return downloadRequest{fileid: fid}, nil
}

func encodeDownloadResponse(w http.ResponseWriter, response interface{}) error {
	resp := response.(downloadResponse)
	w.Header().Set("etag", "whatisetag")
	io.Copy(w, resp.filecontent)
	return nil
}
