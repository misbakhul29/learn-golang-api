package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Protocols string
	Host      string
	Port      string

	HostDB     string
	PortDB     string
	UserDB     string
	PasswordDB string
	NameDB     string

	ApiKey            string
	ApiRateLimit      float64
	ApiRateLimitBurst int
}

func LoadEnv() *Config {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	rateLimit, err := strconv.Atoi(os.Getenv("API_RATE_LIMIT"))
	if err != nil {
		log.Fatalf("Error parsing API_RATE_LIMIT: %v", err)
	}

	burst, err := strconv.Atoi(os.Getenv("API_RATE_LIMIT_BURST"))
	if err != nil {
		log.Fatalf("Error parsing API_RATE_LIMIT_BURST: %v", err)
	}

	cfg.Protocols = os.Getenv("PROTOCOLS")
	cfg.Host = os.Getenv("HOST")
	cfg.Port = os.Getenv("PORT")

	cfg.HostDB = os.Getenv("DB_HOST")
	cfg.PortDB = os.Getenv("DB_PORT")
	cfg.UserDB = os.Getenv("DB_USER")
	cfg.PasswordDB = os.Getenv("DB_PASSWORD")
	cfg.NameDB = os.Getenv("DB_NAME")

	cfg.ApiKey = os.Getenv("API_KEY")
	cfg.ApiRateLimit = float64(rateLimit)
	cfg.ApiRateLimitBurst = burst

	return &cfg
}

func (c *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s://%s:%s", c.Protocols, c.Host, c.Port)
}
