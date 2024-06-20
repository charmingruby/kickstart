package integration

import (
	"bytes"
	"encoding/json"
	"io"
)

func writeBody(body []byte) *bytes.Reader {
	return bytes.NewReader(body)
}

func parseRequest[T any](r *T, body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(&r); err != nil {
		return err
	}

	return nil
}
