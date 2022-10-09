package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/bytepass/server/config"
	"github.com/bytepass/server/crypto"
)

func SendEmailActivationKey(email string, userId string) error {
	token, err := crypto.GenerateActivationJwt(userId)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/api/user/verifyEmail?token=%s", config.Config.Api.Domain, token)

	subject := "Activate your BytePass account!"

	templateData := struct {
		URL string
	}{
		URL: url,
	}
	body, err := parseTemplate(config.EmailActivationTemplate, templateData)
	if err != nil {
		return err
	}

	err = sendEmail(email, subject, *body)
	if err != nil {
		return err
	}

	return nil
}

func sendEmail(email string, subject string, body string) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msgString := fmt.Sprintf("To: %s\nSubject: %s\n%s\n%s", email, subject, mime, body)

	msg := []byte(msgString)

	auth := smtp.PlainAuth("", config.SmtpUser, config.SmtpPassword, config.SmtpHost)

	// convert email string to []string
	var emails []string
	emails = append(emails, email)

	err := smtp.SendMail(config.SmtpAddress, auth, config.SmtpUser, emails, msg)

	return err
}

func parseTemplate(templateData string, data interface{}) (*string, error) {
	t, err := template.New("email.html").Parse(templateData)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return nil, err
	}

	body := buf.String()

	return &body, nil
}
