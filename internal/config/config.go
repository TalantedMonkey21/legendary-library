package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
	DB   DBConfig
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	return &Config{
		Port: GetEnv("PORT", ":8080"),
		DB:   LoadDBConfig(),
	}
}

func GetEnv(key, defValue string) string {
	if key == "" {
		log.Printf("Not found %v, use default\n", key)
		return defValue
	}
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Not found %v value, use default\n", key)
		return defValue
	}
	return value
}
