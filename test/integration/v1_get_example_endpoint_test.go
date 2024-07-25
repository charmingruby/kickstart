package integration

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/kickstart/internal/common/core"
	"github.com/charmingruby/kickstart/internal/example/transport/rest/endpoint/example_endpoint_v1"
	"github.com/charmingruby/kickstart/test/factory"
	"github.com/charmingruby/kickstart/test/integration/helper"
)

func (s *Suite) Test_GetExampleEndpoint() {
	s.Run("it should be able get an example by id", func() {
		name := "example"

		example, err := factory.MakeExample(s.exampleRepo, name)
		s.NoError(err)

		route := s.V1Route(
			fmt.Sprintf("/examples/%s", example.ID),
		)

		res, err := http.Get(route)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusOK, res.StatusCode)

		data := example_endpoint_v1.GetExampleResponse{}
		err = helper.ParseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal("example found", data.Message)
		s.Equal(example.ID, data.Data.ID)
		s.Equal(example.Name, data.Data.Name)
	})

	s.Run("it should be not able get an example by a nonexistent id", func() {
		route := s.V1Route(
			fmt.Sprintf("/examples/%s", "invalid id"),
		)

		res, err := http.Get(route)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusNotFound, res.StatusCode)

		data := example_endpoint_v1.GetExampleResponse{}
		err = helper.ParseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal(core.NewNotFoundErr("example").Error(), data.Message)
		s.Equal(http.StatusNotFound, data.Code)
	})
}
