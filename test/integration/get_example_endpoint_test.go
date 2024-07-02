package integration

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/kickstart/internal/core"
	"github.com/charmingruby/kickstart/internal/domain/example/example_entity"
	v1 "github.com/charmingruby/kickstart/internal/infra/transport/rest/endpoint/v1"
)

func (s *Suite) Test_GetExampleEndpoint() {
	s.Run("it should be able get an example by id", func() {
		example, err := example_entity.NewExample("Dummy Name")
		s.NoError(err)

		err = s.exampleRepo.Store(example)
		s.NoError(err)

		route := fmt.Sprintf("/v1/examples/%s", example.ID)

		res, err := http.Get(s.Route(route))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusOK, res.StatusCode)

		data := v1.GetExampleResponse{}
		err = parseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal("example found", data.Message)
		s.Equal(example.ID, data.Data.ID)
		s.Equal(example.Name, data.Data.Name)
	})

	s.Run("it should be not able get an example by a nonexistent id", func() {
		route := fmt.Sprintf("/v1/examples/%s", "invalid_id")

		res, err := http.Get(s.Route(route))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusNotFound, res.StatusCode)

		data := v1.GetExampleResponse{}
		err = parseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal(core.NewNotFoundErr("example").Error(), data.Message)
		s.Equal(http.StatusNotFound, data.Code)
	})
}
