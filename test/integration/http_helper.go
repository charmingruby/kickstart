package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func writeBody(body []byte) *bytes.Reader {
	return bytes.NewReader(body)
}

func readBody(res *http.Response) ([]byte, error) {
	return io.ReadAll(res.Body)
}

func parseRequest[T any](r *T, body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(&r); err != nil {
		return err
	}

	return nil
}

type errorResponse struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
