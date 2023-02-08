package db

import (
	model "PharmaProject/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	conn, err := LoadConfig()
	dsn := "dbname=" + conn.DBName + " host=" + conn.DBHost + " user=" + conn.DBUser + " password=" + conn.DBPassword + " port=" + conn.DBPort + " sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(model.User{}, model.Cart{}, model.Medicine{}, model.Order{})
	return db
}
