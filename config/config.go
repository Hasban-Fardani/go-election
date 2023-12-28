package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	JwtSecret       string
	JwtExpireMinute int
	DB              DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

var Data *Config

func init() {
	data, _ := LoadConfig()
	Data = data
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")

	minute, _ := strconv.Atoi(getEnv("JWT_EXPIRE_MINUTE", "10"))

	return &Config{
		Port:            getEnv("APP_PORT", "3000"),
		JwtSecret:       getEnv("JWT_SECRET", "secret"),
		JwtExpireMinute: minute,
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnv("DB_PORT", "3306"),
			Username: getEnv("DB_USERNAME", "root"),
			Password: getEnv("DB_PASSWORD", ""),
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
