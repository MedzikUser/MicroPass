package ciphers

import (
	"net/http"

	"github.com/MedzikUser/MicroPass/database"
	"github.com/MedzikUser/MicroPass/server/api/auth"
	"github.com/MedzikUser/MicroPass/server/errors"
	"github.com/gin-gonic/gin"
)

func delete(ctx *gin.Context) {
	token := auth.GetAccessToken(ctx)
	if token == nil {
		return
	}

	id := ctx.Param("id")

	cipher := database.Cipher{
		Id:    id,
		Owner: token.User.Id,
	}
	err := cipher.Delete()
	if err != nil {
		errors.ErrDatabase(ctx)
		return
	}

	ctx.Status(http.StatusOK)
}
