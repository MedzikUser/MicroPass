package ciphers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/MedzikUser/MicroPass/server/api/auth"
	"github.com/MedzikUser/MicroPass/server/errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func list(ctx *gin.Context) {
	token := auth.GetAccessToken(ctx)
	if token == nil {
		return
	}

	var cipherList []string

	ciphers, err := token.User.TakeOwnedCiphers()
	if err != nil {
		// if record not found, return empty array
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, gin.H{"ciphers": cipherList})
			return
		}

		errors.ErrDatabase(ctx)
		return
	}

	lastSync := ctx.Query("lastSync")

	for _, cipher := range ciphers {
		if lastSync != "" {
			lastSyncInt, err := strconv.ParseInt(lastSync, 10, 64)
			if err != nil {
				errors.ErrInvalidRequest(ctx)
				return
			}

			lastSyncTime := time.Unix(lastSyncInt, 0)

			if cipher.UpdatedAt.After(lastSyncTime) {
				cipherList = append(cipherList, cipher.Id)
			}
		} else {
			cipherList = append(cipherList, cipher.Id)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ciphers": cipherList,
	})
}
