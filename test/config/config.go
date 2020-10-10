package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config help running test on different enviroments
type Config struct {
	BaseURL string `envconfig:"BASE_URL" required:"true"`
}

// ProcessConfig create new config from enviroment variables
func ProcessConfig() *Config {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("BASE_URL: %v\n", c.BaseURL)
	return &c
}
