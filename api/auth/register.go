package auth

import (
	"net/http"

	"github.com/bytepass/server/database"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var post registerPost

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
	_, err = database.NewUser(post.Email, post.MasterPassword, post.MasterPasswordHint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Email is already taken",
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
