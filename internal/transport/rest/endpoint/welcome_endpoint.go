package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type welcomeResponse struct {
	Message string `json:"message"`
}

func welcomeEndpoint(c *gin.Context) {
	res := welcomeResponse{
		Message: "Ok",
	}
	c.JSON(http.StatusOK, res)
}
