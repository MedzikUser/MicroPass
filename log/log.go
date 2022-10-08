package log

import (
	"fmt"
	"log"

	"github.com/bytepass/server/config"
)

// Colors
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Magenta = "\033[35m"
	Gray    = "\033[30m"
)

func Info(format string, v ...any) {
	log.Printf(fmt.Sprintf("%s[INFO]%s %s", Magenta, Reset, format), v...)
}

func Error(format string, v ...any) {
	log.Printf(fmt.Sprintf("%s[ERROR]%s %s", Red, Reset, format), v...)
}

func Fatal(format string, v ...any) {
	log.Fatalf(fmt.Sprintf("%s[FATAL]%s %s", Red, Reset, format), v...)
}

func Debug(format string, v ...any) {
	if config.DevelopmentBuild {
		log.Printf(fmt.Sprintf("%s[DEBUG]%s %s", Gray, Reset, format), v...)
	}
}
