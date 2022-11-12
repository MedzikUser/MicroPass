package ciphers

import "github.com/gin-gonic/gin"

func Apply(router *gin.RouterGroup) {
	ciphers := router.Group("/ciphers")

	ciphers.GET("/get/:id", get)
	ciphers.DELETE("/delete/:id", delete)
	ciphers.POST("/insert", insert)
	ciphers.GET("/list", list)
	ciphers.PATCH("/update", update)
}
