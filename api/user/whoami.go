package user

import (
	"fmt"
	"net/http"

	"github.com/bytepass/server/api/utils"
	"github.com/bytepass/server/database"
	"github.com/gin-gonic/gin"
)

func whoami(c *gin.Context) {
	token, err := utils.GetToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("Failed to extract token from headers: %v", err),
		})

		return
	}

	// take user from database
	user, err := database.TakeUserID(token.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "User not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"email":   user.Email,
	})
}
