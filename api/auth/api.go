package auth

import (
	"github.com/gin-gonic/gin"
)

func Apply(r *gin.Engine) {
	auth := r.Group("/api/auth")

	auth.POST("/login", login)
	auth.POST("/register", register)
}
