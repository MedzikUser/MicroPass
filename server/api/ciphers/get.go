package ciphers

import (
	"net/http"

	"github.com/MedzikUser/MicroPass/database"
	"github.com/MedzikUser/MicroPass/server/api/auth"
	"github.com/MedzikUser/MicroPass/server/errors"
	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	token := auth.GetAccessToken(ctx)
	if token == nil {
		return
	}

	id := ctx.Param("id")

	cipher := database.Cipher{
		Id:    id,
		Owner: token.User.Id,
	}
	err := cipher.Take()
	if err != nil {
		errors.ErrDatabase(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       cipher.Id,
		"favorite": cipher.Favorite,
		"data":     cipher.Data,
		"created":  cipher.Model.CreatedAt.Unix(),
		"updated":  cipher.Model.UpdatedAt.Unix(),
	})
}
