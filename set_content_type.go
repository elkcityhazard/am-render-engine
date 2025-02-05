package amrenderengine

import (
	"bytes"
	"io"
	"net/http"
)

func (tmpCol *TemplateCollection) SetContentType(buf *bytes.Buffer) (string, error) {
	// copy the buf to a new tmp buf to detect content type
	tmpBuf := *buf
	hdrBuf := make([]byte, 512)
	_, err := io.ReadFull(&tmpBuf, hdrBuf)
	if err != nil {

		return "", err
	}
	mimetype := http.DetectContentType(hdrBuf)
	return mimetype, nil
}
