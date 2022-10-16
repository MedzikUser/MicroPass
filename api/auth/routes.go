package auth

import "github.com/gin-gonic/gin"

func Apply(r *gin.Engine) {
	auth := r.Group("/api/auth")

	// 1. TODO: add rate limit for disable brute force login.
	// 2. TODO: implement two factor authentication.

	auth.POST("/login", login)
	auth.POST("/register", register)
	auth.GET("/refresh", refresh)
}
