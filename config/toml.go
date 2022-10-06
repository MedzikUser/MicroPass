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
	PublicKey  string
	PrivateKey string
	Expires    time.Duration
}

type apiConfig struct {
	Address  string
	Tls      bool
	CertFile string
	KeyFile  string
}

var Config config

func init() {
	ParseConfig("config.toml")

	pub, err := ioutil.ReadFile(Config.Jwt.PublicKey)
	if err != nil {
		log.Fatal("Failed to open jwt public key file: %v", err)
	}

	private, err := ioutil.ReadFile(Config.Jwt.PrivateKey)
	if err != nil {
		log.Fatal("Failed to open jwt private key file: %v", err)
	}

	Config.Jwt.PublicKey = string(pub)
	Config.Jwt.PrivateKey = string(private)

	println("Initializing config")
}

func ParseConfig(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Could not open config file: ", err)
	}

	config := &config{}

	err = toml.Unmarshal(data, config)
	if err != nil {
		log.Fatal("Could not parse config file: ", err)
	}

	Config = *config
}
