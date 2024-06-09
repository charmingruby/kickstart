package integration

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/kickstart/internal/domain/example"
	"github.com/charmingruby/kickstart/internal/validation"
)

func (s *Suite) Test_GetExampleEndpoint() {
	type getExampleResponse struct {
		Code    int              `json:"status_code"`
		Message string           `json:"message"`
		Data    *example.Example `json:"data,omitempty"`
	}

	s.Run("it should be able get an example by id", func() {
		example, err := example.NewExample("Dummy Name")
		s.NoError(err)

		err = s.exampleRepo.Store(example)
		s.NoError(err)

		route := fmt.Sprintf("/v1/examples/%s", example.ID)
		res, err := http.Get(s.Route(route))
		s.NoError(err)
		s.Equal(http.StatusOK, res.StatusCode)
		defer res.Body.Close()

		data := getExampleResponse{}
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
		s.Equal(http.StatusNotFound, res.StatusCode)
		defer res.Body.Close()

		data := getExampleResponse{}
		err = parseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal(validation.NewNotFoundErr("example").Error(), data.Message)
		s.Equal(http.StatusNotFound, data.Code)
	})
}
