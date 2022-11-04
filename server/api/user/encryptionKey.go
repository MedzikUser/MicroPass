package user

import (
	"net/http"

	"github.com/MedzikUser/AwesomeVault/server/api/auth"
	"github.com/gin-gonic/gin"
)

func encryptionKey(ctx *gin.Context) {
	token := auth.GetAccessToken(ctx)
	if token == nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"encryption_key": token.User.EncryptionKey,
	})
}
