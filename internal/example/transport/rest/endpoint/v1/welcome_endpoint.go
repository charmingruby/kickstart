package v1

import (
	"github.com/charmingruby/kickstart/internal/common/api/rest"
	"github.com/gin-gonic/gin"
)

// Welcome godoc
//
//	@Summary		Health Check
//	@Description	Health Check
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	rest.Response
//	@Router			/welcome [get]
func welcomeEndpoint(c *gin.Context) {
	rest.NewOkResponse(c, "OK!", nil)
}
