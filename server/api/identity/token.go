package identity

import (
	"net/http"

	"github.com/MedzikUser/MicroPass/database"
	"github.com/MedzikUser/MicroPass/server/errors"
	"github.com/MedzikUser/MicroPass/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func token(ctx *gin.Context) {
	var formData ConnectData
	ctx.Bind(&formData)

	// * Refresh token authentication
	if formData.GrantType == GrantTypeRefreshToken {
		refreshToken := formData.RefreshToken

		if refreshToken == nil {
			errors.ErrInvalidRequest(ctx)
			return
		}

		// validate refresh token
		userId, err := utils.ValidateRefreshToken(*refreshToken)
		if err != nil {
			errors.ErrInvalidToken(ctx)
			return
		}

		// create user model
		user := &database.User{
			Id: *userId,
		}
		// take the user from the database
		err = user.Take()
		if err != nil {
			// record not found
			if errors.Is(err, gorm.ErrRecordNotFound) {
				errors.ErrUserNotFound(ctx)
				return
			}

			// other error
			errors.ErrDatabase(ctx)
			return
		}

		// generate access token
		accessToken, err := user.GenerateAccessToken()
		if err != nil {
			// record not found
			if errors.Is(err, gorm.ErrRecordNotFound) {
				errors.ErrUserNotFound(ctx)
				return
			}

			// other error
			errors.ErrGenerateToken(ctx)
			return
		}

		// set response headers
		ctx.Writer.Header().Set("Access-Token", accessToken)

		// send response body
		ctx.JSON(http.StatusOK, gin.H{
			"access_token": accessToken,
		})

		return
	}

	// * Password authentication
	if formData.GrantType == GrantTypePassword {
		// TODO implement 2FA (two factor authentication).

		email := formData.Email
		password := formData.Password

		if email == nil || len(*email) == 0 || password == nil || len(*password) == 0 {
			errors.ErrInvalidRequest(ctx)
			return
		}

		// create user model
		user := &database.User{
			Email:    *email,
			Password: *password,
		}
		// take the user from the database
		err := user.Take()
		if err != nil {
			// record not found
			if errors.Is(err, gorm.ErrRecordNotFound) {
				errors.ErrUserNotFound(ctx)
				return
			}

			// password missmatch
			if errors.Is(err, database.ErrPasswordMismatch) {
				errors.ErrPasswordMismatch(ctx)
				return
			}

			// other error
			errors.ErrDatabase(ctx)
			return
		}

		// TODO add email verification

		// generate access token
		accessToken, err := user.GenerateAccessToken()
		if err != nil {
			errors.ErrGenerateToken(ctx)
			return
		}

		// generate refresh token
		refreshToken, err := user.GenerateRefreshToken()
		if err != nil {
			errors.ErrGenerateToken(ctx)
			return
		}

		// set response headers
		ctx.Writer.Header().Set("Access-Token", accessToken)
		ctx.Writer.Header().Set("Refresh-Token", refreshToken)

		// send response body
		ctx.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})

		return
	}

	// unknown grant type
	errors.ErrInvalidGrantType(ctx)
}

var (
	GrantTypeRefreshToken = "refresh_token"
	GrantTypePassword     = "password"
)

type ConnectData struct {
	GrantType string `form:"grant_type" json:"grant_type" binding:"required"` // refresh_token or password

	// Needed for grant_type="refresh_token"
	RefreshToken *string `form:"refresh_token" json:"refresh_token"`

	// Needed for grant_type="password"
	Email    *string `form:"email"    json:"email"`
	Password *string `form:"password" json:"password"`
}
