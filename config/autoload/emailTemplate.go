package autoload

import (
	"io/ioutil"

	"github.com/bytepass/server/config"
	"github.com/bytepass/server/log"
)

func init() {
	log.Debug("Loading email templates")

	activationTemplate, err := ioutil.ReadFile("assets/email/activation.html")
	if err != nil {
		log.Fatal("Failed to read activation email template: %v", err)
	}

	config.EmailActivationTemplate = string(activationTemplate)
}
