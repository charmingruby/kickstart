package integration

import (
	"encoding/json"
	"net/http"

	"github.com/charmingruby/kickstart/internal/common/api/rest"
	"github.com/charmingruby/kickstart/internal/common/core/validation"
	v1 "github.com/charmingruby/kickstart/internal/example/transport/rest/endpoint/v1"
	"github.com/charmingruby/kickstart/test/integration/helper"
)

func (s *Suite) Test_CreateExampleEndpoint() {
	s.Run("it should be able to create an example", func() {
		route := s.V1Route(
			"/examples",
		)

		payload := v1.CreateExampleRequest{Name: "Dummy name"}
		body, err := json.Marshal(payload)
		s.NoError(err)

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusCreated, res.StatusCode)

		data := rest.Response{}
		err = helper.ParseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal("example created successfully", data.Message)
	})

	s.Run("it should be not able to create an invalid example", func() {
		route := s.V1Route(
			"/examples",
		)

		payload := v1.CreateExampleRequest{Name: "12"}
		body, err := json.Marshal(payload)
		s.NoError(err)

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		s.Equal(http.StatusUnprocessableEntity, res.StatusCode)
		defer res.Body.Close()

		data := rest.Response{}
		err = helper.ParseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal(validation.ErrMinLength("name", "3"), data.Message)
	})
}
