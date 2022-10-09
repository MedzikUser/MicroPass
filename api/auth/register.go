package auth

import (
	"net/http"
	"strings"

	"github.com/bytepass/server/api/utils"
	"github.com/bytepass/server/database"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var post registerPost

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

	// inset user into database
	user, err := database.NewUser(post.Email, post.MasterPassword, post.MasterPasswordHint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Email is already taken",
		})

		return
	}

	// send activation email
	err = utils.SendEmailActivationKey(user.Email, user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to send activation email",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type registerPost struct {
	Email              string
	MasterPassword     string
	MasterPasswordHint string
}
