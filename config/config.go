package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DB   DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("../.env")

	return &Config{
		Port: getEnv("APP_PORT", "3000"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			Username: getEnv("DB_USERNAME", "root"),
			Password: getEnv("DB_PASSWORD", "password"),
			Database: getEnv("DB_NAME", "go_election"),
		},
	}, err
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
