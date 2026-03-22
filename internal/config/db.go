package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database")
	}
	fmt.Println("Connected to database", DB)
	
	err = DB.AutoMigrate(&Note{})
	if err != nil {
		fmt.Println("Failed to migrate database:", err)
	}

	fmt.Println("Database migrated successfully")
}