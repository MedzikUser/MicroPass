package utils

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	var config zap.Config

	// if the env variable `MICROPASS_DEBUG` is set to true, use development logger config
	if os.Getenv("MICROPASS_DEBUG") == "true" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.OutputPaths = []string{"server.log"}

	var err error
	Log, err = config.Build()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
}
