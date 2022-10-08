package auth

import (
	"net/http"

	"github.com/bytepass/server/crypto"
	"github.com/bytepass/server/database"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var post loginPost

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

	// generate access token
	accessToken, err := crypto.GenerateJWT(user.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Failed to generate access token",
		})

		return
	}

	refreshToken := "none"

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
