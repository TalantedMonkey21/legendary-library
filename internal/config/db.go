package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:     GetEnv("POSTGRES_HOST", "localhost"),
		User:     GetEnv("POSTGRES_USER", "admin"),
		Password: GetEnv("POSTGRES_PASSWORD", "supersecret"),
		Name:     GetEnv("POSTGRES_DB", "lectures"),
		Port:     GetEnv("POSTGRES_PORT", "5432"),
		SSLMode:  GetEnv("POSTGRES_SSLMODE", "disable"),
	}
}

func DSN(db DBConfig) string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		db.Host,
		db.User,
		db.Password,
		db.Name,
		db.Port,
		db.SSLMode,
	)
}

func ConnectDB(cfg DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DSN(cfg)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
