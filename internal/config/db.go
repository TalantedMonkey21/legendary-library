package config

import (
	"fmt"

	"github.com/TalantedMonkey21/GoLectures/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDb() models.DbConfig {
	var d models.DbConfig
	d = models.DbConfig{
		Host: GetEnv("POSTGRES_HOST", "localhost"),
		User: GetEnv("POSTGRES_USER", "admin"),		
		Password: GetEnv("POSTGRES_PASSWORD", "supersecret"),
		Dbname: GetEnv("POSTGRES_DB", "lectures"),
		Port: GetEnv("POSTGRES_PORT", "5432"),
		Sslmode: GetEnv("POSTGRES_SSLMODE", "disable"),
	}
	return d
}

func Dsn(db models.DbConfig) string {
	return 	fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		db.Host,
		db.User,
		db.Password,
		db.Dbname,
		db.Port,
		db.Sslmode,
	)
}

func ConnectDb(cfg models.DbConfig) (*gorm.DB, error) {
	connect, err := gorm.Open(postgres.Open(Dsn(cfg)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = connect.AutoMigrate(&models.Note{})
	if err != nil {
		return nil, err
	}
	return connect, nil
}