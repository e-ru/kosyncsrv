package config

import (
	"kosyncsrv/database"

	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
)

type Configuration struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"127.0.0.1"`
	Port          string `env:"PORT" envDefault:"8080"`
	Ssl           bool   `env:"SSL" envDefault:"false"`
	SslCert       string `env:"SSL_CERT" envDefault:"./cert.pem"`
	SslPrivKey    string `env:"SSL_PRIV_KEY" envDefault:"./cert.key"`
	DBDriverName  string `env:"DB_DRIVER_NAME" envDefault:"sqlite3"`
	DBFileName    string `env:"DB_FILE_NAME" envDefault:"syncdata.db"`
}

func NewConfiguration() (*Configuration, error) {
	cfg := &Configuration{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

var DbHandlerFromConfig = func(config Configuration) (database.DBHandler, error) {
	db, err := sqlx.Connect(config.DBDriverName, config.DBFileName)
	if err != nil {
		return nil, err
	}

	return database.NewDBHandler(db), nil
}
