package user

import (
	"net/http"

	"github.com/MedzikUser/MicroPass/server/api/auth"
	"github.com/gin-gonic/gin"
)

func whoami(ctx *gin.Context) {
	token := auth.GetAccessToken(ctx)
	if token == nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       token.User.Id,
		"email":    token.User.Email,
		"username": token.User.Username,
	})
}
