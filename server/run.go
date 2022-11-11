package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MedzikUser/MicroPass/server/api"
	"github.com/MedzikUser/MicroPass/server/errors"
	"github.com/MedzikUser/MicroPass/utils"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go.uber.org/zap/zapcore"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// add zap logger middleware
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		utils.Log.With(zapcore.Field{
			Key:    "client_ip",
			Type:   zapcore.StringType,
			String: param.ClientIP,
		}, zapcore.Field{
			Key:    "error_message",
			Type:   zapcore.StringType,
			String: param.ErrorMessage,
		}, zapcore.Field{
			Key:     "status",
			Type:    zapcore.Int64Type,
			Integer: int64(param.StatusCode),
		}, zapcore.Field{
			Key:    "latency",
			Type:   zapcore.StringType,
			String: param.Latency.String(),
		}, zapcore.Field{
			Key:    "request",
			Type:   zapcore.StringType,
			String: fmt.Sprintf("%s %s", param.Request.Proto, param.Request.UserAgent()),
		}).Info(fmt.Sprintf("%s %s", param.Method, param.Path))

		return ""
	}))

	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "I am working!")
	})

	// apply api router
	api.Apply(router)

	// 404 handler
	router.NoRoute(func(ctx *gin.Context) {
		errors.ErrPageNotFound(ctx)
	})

	// run https server if enabled and run http redirect to https
	if utils.Config.Https.Enabled {
		// * Run https server

		go func() {
			if utils.Config.Https.Enabled {
				log.Println("Starting https server on", utils.Config.Https.Address)

				log.Fatal(router.RunTLS(utils.Config.Https.Address, utils.Config.Https.CertFile, utils.Config.Https.KeyFile))
			}
		}()

		// * Run http redirect

		router := gin.Default()

		// add https redirect middleware
		router.Use(func(c *gin.Context) {
			secureMiddleware := secure.New(secure.Options{
				SSLRedirect: true,
				SSLHost:     strings.Split(c.Request.Host, ":")[0] + utils.Config.Https.Address,
			})

			err := secureMiddleware.Process(c.Writer, c.Request)
			// if error, do not continue
			if err != nil {
				return
			}

			c.Next()
		})

		if utils.Config.Http.Enabled {
			log.Fatal(router.Run(utils.Config.Http.Address))
		} else {
			<-make(chan int)
		}
	} else {
		log.Println("Starting http server on", utils.Config.Http.Address)

		log.Fatal(router.Run(utils.Config.Http.Address))
	}
}
