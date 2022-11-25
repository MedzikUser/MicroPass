package ciphers

import (
	"net/http"
	"strconv"

	"github.com/MedzikUser/MicroPass/server/api/auth"
	"github.com/MedzikUser/MicroPass/server/errors"
	"github.com/gin-gonic/gin"
)

func list(ctx *gin.Context) {
	token := auth.GetAccessToken(ctx)
	if token == nil {
		return
	}

	lastSyncQuery := ctx.Query("lastSync")
	var lastSync int64
	if lastSyncQuery != "" {
		var err error
		lastSync, err = strconv.ParseInt(lastSyncQuery, 10, 64)
		if err != nil {
			errors.ErrInvalidRequest(ctx)
			return
		}
	}

	ciphers, err := token.User.TakeOwnedCiphers(&lastSync)
	if err != nil {
		errors.ErrDatabase(ctx)
		return
	}

	var updated []string
	var deleted []string

	for _, cipher := range ciphers {
		if cipher.DeletedAt != 0 {
			deleted = append(deleted, cipher.Id)
		} else {
			updated = append(updated, cipher.Id)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"updated": updated,
		"deleted": deleted,
	})
}
