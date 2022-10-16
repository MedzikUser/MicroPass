package smtp

import (
	"fmt"
	"net/smtp"

	"github.com/bytepass/server/config"
)

var mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

func Send(email string, subject string, body string) error {
	msgStr := fmt.Sprintf("To: %s\nSubject: %s\n%s\n%s", email, subject, mime, body)
	msg := []byte(msgStr)

	auth := smtp.PlainAuth("", config.SmtpUser, config.SmtpPassword, config.SmtpHost)

	// convert the email string to a string slice
	var emails []string
	emails = append(emails, email)

	err := smtp.SendMail(config.SmtpAddress, auth, config.SmtpUser, emails, msg)

	return err
}
