package conn

import (
	"PharmaProject/config"
	model "PharmaProject/models"
	"fmt"
	"net/url"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	//
	config := config.DB()
	return Connect(config)
}

func Connect(config *config.Database) *gorm.DB {

	host := fmt.Sprintf("%s:%s", config.Host, config.Port)
	uri := url.URL{
		Scheme: "postgres",
		Host:   host,
		Path:   config.Name,
		User:   url.UserPassword(config.Username, config.Password),
	}
	if config.Options != nil {
		val := url.Values(config.Options)
		uri.RawQuery = val.Encode()
	}
	// fmt.Println(uri.String())
	d, err := gorm.Open("postgres",uri.String())
	if err != nil {
		panic(err)
	}
	db=d
	db.AutoMigrate(model.User{}, model.Cart{}, model.Medicine{}, model.Order{})
	db.Debug()
	return db
}
