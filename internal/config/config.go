package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)
//databasesql
//gorm
type Config struct {
	Port string
	Db DbConfig
}

type DbConfig struct {
	Host string
	User string
	Password string
	Dbname string
	Port string
	Sslmode string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не нашел .env файл(")
	}

	cfg := &Config{
		Port: getEnv("PORT", ":8080"),
		Db: DbConfig{
			getEnv("POSTGRES_HOST", "localhost"),
			getEnv("POSTGRES_USER", "admin"),
			getEnv("POSTGRES_PASSWORD", "supersecret"),
			getEnv("POSTGRES_DB", "lectures"),
			getEnv("POSTGRES_PORT", "5432"),
			getEnv("Sslmode", "disable"),
		},
	}
	return cfg
}

func(db DbConfig) GetDsn() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", db.Host, db.User, db.Password, db.Dbname, db.Port, db.Sslmode)
}

func getEnv(key, defValue string) string {
	if key == "" {
		fmt.Printf("Not found %v, use default\n", key)
		return defValue
	}
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Not found %v value, use default\n", key)
		return defValue
	}
	return value
}