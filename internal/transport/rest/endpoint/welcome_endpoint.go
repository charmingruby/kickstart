package endpoint

import (
	"github.com/gin-gonic/gin"
)

func welcomeEndpoint(c *gin.Context) {
	newOkResponse(c, "OK!", nil)
}
