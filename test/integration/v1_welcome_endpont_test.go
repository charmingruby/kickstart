package integration

import (
	"net/http"

	"github.com/charmingruby/kickstart/test/integration/helper"
)

func (s *Suite) Test_WelcomeEndpoint() {
	type welcomeResponse struct {
		Code    int    `json:"status_code"`
		Message string `json:"message"`
	}

	s.Run("it should be able to get a welcome message", func() {
		route := s.V1Route("/welcome")

		res, err := http.Get(route)
		s.NoError(err)
		s.Equal(http.StatusOK, res.StatusCode)
		defer res.Body.Close()

		data := welcomeResponse{}
		err = helper.ParseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal("OK!", data.Message)
	})
}
