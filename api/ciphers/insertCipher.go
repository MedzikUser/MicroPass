package ciphers

import (
	"fmt"
	"net/http"

	"github.com/bytepass/server/api/utils"
	"github.com/bytepass/server/database"
	"github.com/gin-gonic/gin"
)

func insertCipher(c *gin.Context) {
	// validate user credentials
	token, err := utils.GetToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("Failed to extract token from headers: %v", err),
		})

		return
	}

	// parse request data
	var post insertCipherPost

	err = c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request data",
		})

		return
	}

	// inser cipher to the database
	var cipher database.Cipher
	cipher.UserId = &token.UserId
	cipher.Data = post.Data
	cipher.Collection = post.Collection

	cipher, err = cipher.Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("Failed to insert cipher to database: %v", err),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type insertCipherPost struct {
	Data       string
	Collection *string
}
