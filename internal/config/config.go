package config

import (
	"log"
	"os"

	"github.com/TalantedMonkey21/GoLectures/internal/models"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Db models.DbConfig
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	return &Config{
		Port: GetEnv("PORT", ":8080"),
		Db: ConfigDb(),
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