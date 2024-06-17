package endpoint

import (
	"github.com/gin-gonic/gin"
)

// Welcome godoc
//
//	@Summary		Welcome
//	@Description	Health Check
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	Response
//	@Router			/welcome [get]
func welcomeEndpoint(c *gin.Context) {
	newOkResponse(c, "OK!", nil)
}
