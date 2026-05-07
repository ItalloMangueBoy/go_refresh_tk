package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	DatabaseURL string

	AccessTokenSecret  string
	AccessTokenTTL     time.Duration

	RefreshTokenLength int
	RefreshTokenTTL    time.Duration
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	DatabaseURL = GetEnvString("DATABASE_URL", "")

	AccessTokenSecret = GetEnvString("ACCESS_TOKEN_SECRET", "")
	AccessTokenTTL = getEnvDuration("ACCESS_TOKEN_TTL", 15*time.Minute)

	RefreshTokenLength = GetEnvInt("REFRESH_TOKEN_LENGTH", 32)
	RefreshTokenTTL = getEnvDuration("REFRESH_TOKEN_TTL", 24*time.Hour)

	return nil
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return parsed
}

func GetEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsed
}
