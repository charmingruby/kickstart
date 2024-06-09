package integration

import (
	"net/http"
)

func (s *Suite) Test_WelcomeEndpoint() {
	type welcomeResponse struct {
		Code    int    `json:"status_code"`
		Message string `json:"message"`
	}

	s.Run("it should be able to get a welcome message", func() {
		res, err := http.Get(s.Route("/v1/welcome"))
		s.NoError(err)
		s.Equal(http.StatusOK, res.StatusCode)
		defer res.Body.Close()

		data := welcomeResponse{}
		err = parseRequest(&data, res.Body)
		s.NoError(err)

		s.Equal("OK!", data.Message)
	})
}
