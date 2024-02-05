package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB     Postgres
	DBConn string

	MQ RabbitMQ
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	SSLMode  string
	DBName   string
}

type RabbitMQ struct {
	URI string
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}
	cfg.DBConn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName, cfg.DB.SSLMode)

	if err := envconfig.Process("amqp", &cfg.MQ); err != nil {
		return nil, err
	}

	return cfg, nil
}
