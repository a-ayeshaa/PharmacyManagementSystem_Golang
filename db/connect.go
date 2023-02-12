package db

import (
	model "PharmaProject/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	//
	LoadDB()
	con:= DB()
	// println(con.Name)
	// conn, err := LoadConfig()
	dsn := "dbname=" + con.Name + " host=" + con.Host + " user=" + con.Username + " password=" + con.Password + " port=" + con.Port + " sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(model.User{}, model.Cart{}, model.Medicine{}, model.Order{})
	db.Debug()
	return db
}

