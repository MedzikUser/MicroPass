package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	SmtpHost     = os.Getenv("SMTP_HOST")
	SmtpPort     = os.Getenv("SMTP_PORT")
	SmtpAddress  = fmt.Sprintf("%s:%s", SmtpHost, SmtpPort)
	SmtpUser     = os.Getenv("SMTP_USER")
	SmtpPassword = os.Getenv("SMTP_PASSWORD")
)
