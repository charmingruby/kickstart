package integration

import (
	"encoding/json"
	"net/http"

	"github.com/charmingruby/kickstart/internal/validation"
	"github.com/stretchr/testify/assert"
)

func (s *Suite) Test_CreateExampleEndpoint() {
	type createExamplePayload struct {
		Name string `json:"name"`
	}

	s.Run("it should be able to create an example", func() {
		payload := createExamplePayload{Name: "Dummy name"}
		body, err := json.Marshal(payload)
		assert.NoError(s.T(), err)

		res, err := http.Post(s.Route("/examples"), contentType, writeBody(body))
		s.NoError(err)
		s.Equal(http.StatusCreated, res.StatusCode)
		defer res.Body.Close()

		_, err = readBody(res)
		s.NoError(err)
	})

	s.Run("it should be not able to create an invalid example", func() {
		payload := createExamplePayload{Name: "12"}
		body, err := json.Marshal(payload)
		assert.NoError(s.T(), err)

		res, err := http.Post(s.Route("/examples"), contentType, writeBody(body))
		s.NoError(err)
		s.Equal(http.StatusBadRequest, res.StatusCode)
		defer res.Body.Close()

		data := errorResponse{}
		err = parseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal(validation.ErrMinLength("name", "3"), data.Message)
		s.Equal(http.StatusBadRequest, data.Code)
	})
}
