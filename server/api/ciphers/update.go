package ciphers

import (
	"net/http"

	"github.com/MedzikUser/MicroPass/database"
	"github.com/MedzikUser/MicroPass/server/api/auth"
	"github.com/MedzikUser/MicroPass/server/errors"
	"github.com/gin-gonic/gin"
)

func update(ctx *gin.Context) {
	token := auth.GetAccessToken(ctx)
	if token == nil {
		return
	}

	var formData cipherUpdateData
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

	// update cipher in the database
	err := cipher.Update()
	if err != nil {
		errors.ErrDatabase(ctx)
		return
	}

	ctx.Status(http.StatusOK)
}

type cipherUpdateData struct {
	Id   string `form:"id" json:"id" binding:"required"`
	Data string `form:"data" json:"data" binding:"required"`
}
