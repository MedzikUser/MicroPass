package utils

import (
	"fmt"

	"github.com/bytepass/server/config"
	"github.com/bytepass/server/crypto"
	"github.com/bytepass/server/smtp"
)

func SendEmailActivationKey(email string, userId string) error {
	// generate activation token
	token, err := crypto.GenerateActivationJwt(userId)
	if err != nil {
		return fmt.Errorf("generating activation token error: %v", err)
	}

	// format a account activation url
	url := fmt.Sprintf("%s/api/user/verifyEmail?token=%s", config.Config.Api.Domain, token)

	// email subject
	subject := "Activate your BytePass account!"

	// parse email body
	body, err := smtp.ParseActivationTemplate(url)
	if err != nil {
		return fmt.Errorf("parsing activation email template error: %v", err)
	}

	// send the email
	err = smtp.Send(email, subject, *body)
	if err != nil {
		return fmt.Errorf("error when trying to send email: %v", err)
	}

	return nil
}
