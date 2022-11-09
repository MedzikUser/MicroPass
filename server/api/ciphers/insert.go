package ciphers

import (
	"net/http"

	"github.com/MedzikUser/AwesomeVault/database"
	"github.com/MedzikUser/AwesomeVault/server/api/auth"
	"github.com/MedzikUser/AwesomeVault/server/errors"
	"github.com/gin-gonic/gin"
)

func insert(ctx *gin.Context) {
	token := auth.GetAccessToken(ctx)
	if token == nil {
		return
	}

	var formData cipherInsertData
	ctx.Bind(&formData)

	if len(formData.Data) == 0 {
		errors.ErrInvalidRequest(ctx)
		return
	}

	// create cipher model
	cipher := database.Cipher{
		Owner: token.User.Id,
		Data:  formData.Data,
	}

	// insert cipher to the database
	err := cipher.Insert()
	if err != nil {
		errors.ErrDatabase(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": cipher.Id,
	})
}

type cipherInsertData struct {
	Data string `form:"data" json:"data" binding:"required"`
}