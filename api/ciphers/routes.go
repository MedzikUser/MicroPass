package ciphers

import (
	"github.com/gin-gonic/gin"
)

func Apply(r *gin.Engine) {
	ciphers := r.Group("/api/ciphers")

	ciphers.GET("/get/:uuid", getCipher)
	ciphers.POST("/insert", insertCipher)
}
