package user

import (
	"github.com/gin-gonic/gin"
)

func Apply(r *gin.Engine) {
	user := r.Group("/api/user")

	user.GET("/whoami", whoami)
}
