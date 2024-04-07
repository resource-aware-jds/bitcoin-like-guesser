package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Endpoint     string `envconfig:"ENDPOINT"`
	ExpectedHash string `envconfig:"EXPECTED_HASH"`
	RoundID      string `envconfig:"ROUND_ID"`
	NodeID       string `envconfig:"NODE_ID"`
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Can't load env file")
	}

	var config Config
	envconfig.MustProcess("", &config)
	return config
}
