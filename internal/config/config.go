package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig *Config

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	AppConfig = &Config{
		AppPort:    GetEnv("APP_PORT", "8080"),
		DBHost:     GetEnv("DB_HOST", "localhost"),
		DBPort:     GetEnv("DB_PORT", "5432"),
		DBUser:     GetEnv("DB_USER", "postgres"),
		DBPassword: GetEnv("DB_PASSWORD", "12345678"),
		DBName:     GetEnv("DB_NAME", "demo_db"),
	}
}

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		return defaultValue
	}
}
