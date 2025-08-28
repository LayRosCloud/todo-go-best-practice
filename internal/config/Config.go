package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port        int  `envconfig:"PORT" default:"5432"`
	Host     string  `envconfig:"HOST_DB" default:"localhost"`
	User     string  `envconfig:"USER_DB" default:"admin"`
	Password string  `envconfig:"PASSWORD_DB" default:"admin"`
	DBName   string  `envconfig:"NAME_DB" default:"todo_base"`
	SSLMode  string  `envconfig:"SSLMODE_DB" default:"false"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Warning: .env file not found: %v", err)
		// Можно продолжить работу, так как переменные могут быть в окружении
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	return &cfg, err
}