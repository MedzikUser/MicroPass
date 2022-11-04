package user

import (
	"github.com/gin-gonic/gin"
)

func Apply(router *gin.RouterGroup) {
	ciphers := router.Group("/user")

	ciphers.GET("/encryption_key", encryptionKey)
}
