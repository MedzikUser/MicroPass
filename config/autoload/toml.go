package autoload

import (
	"io/ioutil"

	"github.com/bytepass/server/config"
	"github.com/bytepass/server/log"
)

func init() {
	log.Debug("Initializing config")

	// parse the configuration file
	tomlConfig, err := config.ParseConfig("config.toml")
	if err != nil {
		log.Fatal("Failed to parse configuration file: %v", err)
	}

	// read jwt public key
	public, err := ioutil.ReadFile(tomlConfig.Jwt.PublicKey)
	if err != nil {
		log.Fatal("Failed to open jwt public key file: %v", err)
	}

	// read jwt private key
	private, err := ioutil.ReadFile(tomlConfig.Jwt.PrivateKey)
	if err != nil {
		log.Fatal("Failed to open jwt private key file: %v", err)
	}

	// save jwt public and private key to configuration variable
	tomlConfig.Jwt.PublicKey = string(public)
	tomlConfig.Jwt.PrivateKey = string(private)

	// move a local config variable to a global variable
	config.Config = *tomlConfig
}
