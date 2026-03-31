package config

import (
	"fmt"
	"github.com/TalantedMonkey21/GoLectures/internal/db"
	"github.com/TalantedMonkey21/GoLectures/internal/models"
	"github.com/TalantedMonkey21/GoLectures/internal/response"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Db models.DbConfig
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return &Config{
		Port: response.GetEnv("PORT", ":8080"),
		Db: db.ConfigDb(),
	}
}