package database

import (
	"fmt"

	"github.com/iwachan14736/travios-backend-service/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func connect() {
	config := config.LocalConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error when connecting to DB")
		panic(err)
	}

	fmt.Println("Database connected")
	db = d
}

func GetDB() *gorm.DB {
	if db == nil {
		connect()
	}
	return db
}
