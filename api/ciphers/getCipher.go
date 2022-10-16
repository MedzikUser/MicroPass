package ciphers

import (
	"fmt"
	"net/http"

	"github.com/bytepass/server/api/utils"
	"github.com/bytepass/server/database"
	"github.com/gin-gonic/gin"
)

func getCipher(c *gin.Context) {
	// validate user credentials
	token, err := utils.GetToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("Failed to extract token from headers: %v", err),
		})

		return
	}

	// get cipher uuid from request data
	uuid := c.Param("uuid")

	// get cipher from the database
	var cipher database.Cipher
	cipher.Id = uuid
	cipher.UserId = &token.UserId

	cipher, err = cipher.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("Failed to get cipher: %v", err),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"cipher":  cipher,
	})
}
