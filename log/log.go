package log

import (
	"fmt"
	"log"
)

// Colors
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Magenta = "\033[35m"
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
