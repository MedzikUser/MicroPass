package auth

import (
	"net/http"
	"strings"

	"github.com/bytepass/server/crypto"
	"github.com/bytepass/server/database"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var post loginPost

	post.Email = strings.ToLower(post.Email)

	// parse request data
	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request data",
		})

		return
	}

	// take user from database
	user, err := database.TakeUser(post.Email, post.MasterPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "User not found",
		})

		return
	}

	if !user.EmailVerified {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "User email address is not verified",
		})

		return
	}

	// generate access token
	accessToken, err := crypto.GenerateJwt(user.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Failed to generate access token",
		})

		return
	}

	// generate refresh token
	refreshToken, err := crypto.GenerateRefreshJwt(user.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Failed to generate refresh token",
		})

		return
	}

	// set response headers
	c.Writer.Header().Set("Access-Token", accessToken)
	c.Writer.Header().Set("Refresh-Token", refreshToken)

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

type loginPost struct {
	Email          string
	MasterPassword string
}
