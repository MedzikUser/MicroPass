package api

import (
	"net/http"

	"github.com/bytepass/server/api/auth"
	"github.com/bytepass/server/api/ciphers"
	"github.com/bytepass/server/api/user"
	"github.com/bytepass/server/config"
	"github.com/bytepass/server/log"
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "I am working!")
	})

	// apply routes
	auth.Apply(r)
	ciphers.Apply(r)
	user.Apply(r)

	// if enabled run https server, if not run http server
	if config.Config.Api.Tls {
		log.Info("🚀 Launched on https://localhost%s", config.Config.Api.Address)

		r.RunTLS(config.Config.Api.Address, config.Config.Api.CertFile, config.Config.Api.KeyFile)
	} else {
		log.Info("🚀 Launched on http://localhost%s", config.Config.Api.Address)

		r.Run(config.Config.Api.Address)
	}
}
