package integration

import (
	"encoding/json"
	"net/http"

	"github.com/charmingruby/kickstart/internal/core"
)

func (s *Suite) Test_CreateExampleEndpoint() {
	type createExamplePayload struct {
		Name string `json:"name"`
	}

	s.Run("it should be able to create an example", func() {
		payload := createExamplePayload{Name: "Dummy name"}
		body, err := json.Marshal(payload)
		s.NoError(err)

		res, err := http.Post(s.Route("/v1/examples"), contentType, writeBody(body))
		s.NoError(err)
		s.Equal(http.StatusCreated, res.StatusCode)
		defer res.Body.Close()

		_, err = readBody(res)
		s.NoError(err)
	})

	s.Run("it should be not able to create an invalid example", func() {
		payload := createExamplePayload{Name: "12"}
		body, err := json.Marshal(payload)
		s.NoError(err)

		res, err := http.Post(s.Route("/v1/examples"), contentType, writeBody(body))
		s.NoError(err)
		s.Equal(http.StatusUnprocessableEntity, res.StatusCode)
		defer res.Body.Close()

		data := errorResponse{}
		err = parseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal(core.ErrMinLength("name", "3"), data.Message)
		s.Equal(http.StatusUnprocessableEntity, data.Code)
	})
}
