package config

import (
	"github.com/caarlos0/env"
)

type Configuration struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"127.0.0.1"`
	Port          string `env:"PORT" envDefault:"8080"`
	SSL           bool   `env:"SSL" envDefault:"false"`
	SSL_CERT      string `env:"SSL_CERT" envDefault:"./cert.pem"`
	SSL_PRIV_KEY  string `env:"SSL_PRIV_KEY" envDefault:"./cert.key"`
	DB_FILE_NAME  string `env:"DB_FILE_NAME" envDefault:"syncdata.db"`
}

func NewConfiguration() (*Configuration, error) {
	cfg := &Configuration{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
