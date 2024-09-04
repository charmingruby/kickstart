package integration

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/kickstart/internal/common/core/custom_err"
	v1 "github.com/charmingruby/kickstart/internal/example/transport/rest/endpoint/v1"
	"github.com/charmingruby/kickstart/test/factory"
	"github.com/charmingruby/kickstart/test/integration/helper"
)

func (s *Suite) Test_GetExampleEndpoint() {
	s.Run("it should be able get an example by id", func() {
		name := "example"

		example, err := factory.MakeExample(s.exampleRepo, name)
		s.NoError(err)

		route := s.V1Route(
			fmt.Sprintf("/examples/%s", example.GetID()),
		)

		res, err := http.Get(route)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusOK, res.StatusCode)

		fmt.Printf("%v", res.Body)

		data := v1.GetExampleResponse{}
		err = helper.ParseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal("example found", data.Message)
		s.Equal(example.GetID(), data.Data.ID)
		s.Equal(example.GetName(), data.Data.Name)
	})

	s.Run("it should be not able get an example by a nonexistent id", func() {
		route := s.V1Route(
			fmt.Sprintf("/examples/%s", "invalid id"),
		)

		res, err := http.Get(route)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusNotFound, res.StatusCode)

		data := v1.GetExampleResponse{}
		err = helper.ParseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal(custom_err.NewNotFoundErr("example").Error(), data.Message)
	})
}
