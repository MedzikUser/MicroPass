package user

import (
	"net/http"

	"github.com/bytepass/server/crypto"
	"github.com/bytepass/server/database"
	"github.com/gin-gonic/gin"
)

func verifyEmail(c *gin.Context) {
	token, exists := c.GetQuery("token")
	if !exists || len(token) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Token parameter is required, but was not provided.",
		})

		return
	}

	userId, err := crypto.ValidateActivationJwt(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to validate activation token",
		})

		return
	}

	var user database.User
	user.Id = *userId

	user, err = user.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to take user from database",
		})

		return
	}

	if user.EmailVerified {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "User is already verified",
		})

		return
	}

	updateUser := database.User{EmailVerified: true}

	err = updateUser.Update(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update user in database",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
