package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PortDb      int  `envconfig:"PORT_DB" default:"5432"`
	Host     string  `envconfig:"HOST_DB" default:"localhost"`
	User     string  `envconfig:"USER_DB" default:"admin"`
	Password string  `envconfig:"PASSWORD_DB" default:"admin"`
	DBName   string  `envconfig:"NAME_DB" default:"todo_base"`
	SSLMode  string  `envconfig:"SSLMODE_DB" default:"false"`
	AccessSecret  string  `envconfig:"ACCESS_SECRET" default:"fdsf32rfvdfv3wfasfh4sadsadcsdgdsfgfdasdas"`
	RefreshSecret  string  `envconfig:"REFRESH_SECRET" default:"hjfdbhjgerhyuighfusadasdasgfadsgvadsfsadf"`
	MigrationPath  string  `envconfig:"MIGRATION_PATH" default:"hjfdbhjgerhyuighfusadasdasgfadsgvadsfsadf"`
}

func LoadConfig() (*Config, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load("../../.env")
		if err != nil {
			log.Printf("Warning: .env file not found: %v", err)
		}
	}

	var cfg Config
	err := envconfig.Process("", &cfg)
	return &cfg, err
}