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
	BaseURL   string

	HostDB     string
	PortDB     string
	UserDB     string
	PasswordDB string
	NameDB     string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
	RedisProtocol int

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

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Error parsing REDIS_DB: %v", err)
	}

	redisProtocol, err := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))
	if err != nil {
		log.Fatalf("Error parsing REDIS_PROTOCOL: %v", err)
	}

	cfg.Protocols = os.Getenv("PROTOCOLS")
	cfg.Host = os.Getenv("HOST")
	cfg.Port = os.Getenv("PORT")
	cfg.BaseURL = os.Getenv("BASE_URL")

	cfg.HostDB = os.Getenv("DB_HOST")
	cfg.PortDB = os.Getenv("DB_PORT")
	cfg.UserDB = os.Getenv("DB_USER")
	cfg.PasswordDB = os.Getenv("DB_PASSWORD")
	cfg.NameDB = os.Getenv("DB_NAME")

	cfg.RedisHost = os.Getenv("REDIS_HOST")
	cfg.RedisPort = os.Getenv("REDIS_PORT")
	cfg.RedisPassword = os.Getenv("REDIS_PASSWORD")
	cfg.RedisDB = redisDB
	cfg.RedisProtocol = redisProtocol

	cfg.ApiKey = os.Getenv("API_KEY")
	cfg.ApiRateLimit = float64(rateLimit)
	cfg.ApiRateLimitBurst = burst

	return &cfg
}

func (c *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *Config) Address() string {
	if c.BaseURL != "" {
		return c.BaseURL
	}
	return fmt.Sprintf("%s://%s:%s", c.Protocols, c.Host, c.Port)
}
