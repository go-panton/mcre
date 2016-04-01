package files

import "io"

type DownloadRequest struct {
	FileId string
}

type DownloadResponse struct {
	FileContent io.Reader
}
