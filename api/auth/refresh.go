package auth

import (
	"fmt"
	"net/http"

	"github.com/bytepass/server/api/utils"
	"github.com/bytepass/server/crypto"
	"github.com/gin-gonic/gin"
)

func refresh(c *gin.Context) {
	token, err := utils.GetRefreshToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("Failed to extract token from headers: %v", err),
		})

		return
	}

	// generate access token
	accessToken, err := crypto.GenerateJwt(token.UserId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Failed to generate access token",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"accessToken": accessToken,
	})
}
