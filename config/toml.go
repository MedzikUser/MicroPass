package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/BurntSushi/toml"
)

type config struct {
	Crypto cryptoConfig
	Jwt    jwtConfig
	Api    apiConfig
	Email  emailConfig
}

type cryptoConfig struct {
	// Length of pseudo-random byte slice (password salt).
	Salt int
	// Number of password iterations.
	Iterations int
}

type jwtConfig struct {
	// The identity of the party who issued the token.
	Issuer string
	// RSA Public Key
	PublicKey string
	// RSA Private Key
	PrivateKey string
	// Access token expiration time in minutes.
	ExpiresAccessToken time.Duration
	// Access token expiration time in minutes.
	ExpiresRefreshToken time.Duration
}

type apiConfig struct {
	Domain   string
	Address  string
	Tls      bool
	CertFile string
	KeyFile  string
}

type emailConfig struct {
	Enabled bool
}

var Config config

func ParseConfig(path string) (*config, error) {
	// read configuration file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}

	var config config

	// unmarshal configuration file
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}

	return &config, nil
}
