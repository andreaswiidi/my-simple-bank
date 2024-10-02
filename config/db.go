package config

import (
	"github.com/andreaswiidi/my-simple-bank/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	dsn := "host=localhost user=root password=123456 dbname=simple_bank port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.AccountBank{}, &models.TransfersHistory{}, &models.TransactionHistory{})

	return db
}
