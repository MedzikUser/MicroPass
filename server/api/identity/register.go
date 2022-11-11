package identity

import (
	"net/http"

	"github.com/MedzikUser/MicroPass/database"
	"github.com/MedzikUser/MicroPass/server/errors"
	"github.com/gin-gonic/gin"
)

func register(ctx *gin.Context) {
	var formData RegisterData
	ctx.Bind(&formData)

	// create user model
	user := &database.User{
		Email:         formData.Email,
		Password:      formData.Password,
		PasswordHint:  formData.PasswordHint,
		EncryptionKey: formData.EncryptionKey,
	}

	// insert user to the database
	err := user.Insert()
	if err != nil {
		// TODO handle duplicate key error

		// other database error
		errors.ErrDatabase(ctx)
		return
	}

	// send empty response
	ctx.Status(http.StatusOK)

	return
}

type RegisterData struct {
	Email         string  `form:"email"          json:"email"          binding:"required"`
	Password      string  `form:"password"       json:"password"       binding:"required"`
	EncryptionKey string  `form:"encryption_key" json:"encryption_key" binding:"required"`
	PasswordHint  *string `form:"password_hint"  json:"password_hint"`
}
