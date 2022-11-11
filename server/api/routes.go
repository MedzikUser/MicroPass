package api

import (
	"github.com/MedzikUser/MicroPass/server/api/ciphers"
	"github.com/MedzikUser/MicroPass/server/api/identity"
	"github.com/MedzikUser/MicroPass/server/api/user"
	"github.com/gin-gonic/gin"
)

func Apply(router *gin.Engine) {
	api := router.Group("/api")

	ciphers.Apply(api)
	identity.Apply(api)
	user.Apply(api)
}
