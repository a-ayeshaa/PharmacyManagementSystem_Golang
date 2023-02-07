package db

import (
	model "PharmaProject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=ayesha password=password dbname=pharmacy_db port=8080 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(model.User{}, model.Cart{}, model.Medicine{}, model.Order{})
	return db
}
