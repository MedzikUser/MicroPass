package config

import (
	"io/ioutil"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/bytepass/server/log"
)

type config struct {
	Crypto cryptoConfig
	Jwt    jwtConfig
	Api    apiConfig
}

type cryptoConfig struct {
	// Length of pseudo-random byte slice (password salt).
	Salt int
	// Number of password iterations.
	Iterations int
}

type jwtConfig struct {
	// RSA Public Key
	PublicKey string
	// RSA Private Key
	PrivateKey string
	// Token expiration time in hours.
	Expires time.Duration
}

type apiConfig struct {
	Address  string
	Tls      bool
	CertFile string
	KeyFile  string
}

var Config config

func init() {
	log.Debug("Initializing config")

	// parse the configuration file
	ParseConfig("config.toml")

	// read jwt public key
	public, err := ioutil.ReadFile(Config.Jwt.PublicKey)
	if err != nil {
		log.Fatal("Failed to open jwt public key file: %v", err)
	}

	// read jwt private key
	private, err := ioutil.ReadFile(Config.Jwt.PrivateKey)
	if err != nil {
		log.Fatal("Failed to open jwt private key file: %v", err)
	}

	// save jwt public and private key to configuration variable
	Config.Jwt.PublicKey = string(public)
	Config.Jwt.PrivateKey = string(private)
}

func ParseConfig(path string) {
	// read configuration file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Could not open config file: %v", err)
	}

	var config config

	// unmarshal configuration file
	err = toml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Could not parse config file: %v", err)
	}

	// move a local config variable to a global variable
	Config = config
}
