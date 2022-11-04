package utils

import (
	"log"

	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{"server.log"}

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	Log = logger
}
