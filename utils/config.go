package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type instanceConfig struct {
	Crypto struct {
		Salt       int
		Iterations int
	}
	Jwt struct {
		Issuer              string
		PublicKey           string `yaml:"public_key"`
		PrivateKey          string `yaml:"private_key"`
		AccessTokenExpires  int64  `yaml:"access_token_expires"`
		RefreshTokenExpires int64  `yaml:"refresh_token_expires"`
	}
	Http struct {
		Enabled bool
		Address string
	}
	Https struct {
		Enabled  bool
		Address  string
		CertFile string `yaml:"cert"`
		KeyFile  string `yaml:"key"`
	}
}

func (c *instanceConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

var Config instanceConfig

func init() {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	if err := Config.Parse(data); err != nil {
		log.Fatal(err)
	}
}
