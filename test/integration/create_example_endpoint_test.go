package integration

import (
	"encoding/json"
	"net/http"

	"github.com/charmingruby/kickstart/internal/core"
	v1 "github.com/charmingruby/kickstart/internal/infra/transport/rest/endpoint/v1"
)

func (s *Suite) Test_CreateExampleEndpoint() {
	s.Run("it should be able to create an example", func() {
		payload := v1.CreateExampleRequest{Name: "Dummy name"}
		body, err := json.Marshal(payload)
		s.NoError(err)

		res, err := http.Post(s.Route("/v1/examples"), contentType, writeBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusCreated, res.StatusCode)

		data := v1.Response{}
		err = parseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal("example created successfully", data.Message)
		s.Equal(http.StatusCreated, data.Code)
	})

	s.Run("it should be not able to create an invalid example", func() {
		payload := v1.CreateExampleRequest{Name: "12"}
		body, err := json.Marshal(payload)
		s.NoError(err)

		res, err := http.Post(s.Route("/v1/examples"), contentType, writeBody(body))
		s.NoError(err)
		s.Equal(http.StatusUnprocessableEntity, res.StatusCode)
		defer res.Body.Close()

		data := v1.Response{}
		err = parseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal(core.ErrMinLength("name", "3"), data.Message)
		s.Equal(http.StatusUnprocessableEntity, data.Code)
	})
}
