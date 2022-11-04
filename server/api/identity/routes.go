package identity

import "github.com/gin-gonic/gin"

func Apply(router *gin.RouterGroup) {
	identity := router.Group("/identity")

	identity.POST("/register", register)
	identity.POST("/token", token)

	// TODO add rate limit for disallow brute force attacks.
	// TODO implement 2FA (two factor authentication).
}
