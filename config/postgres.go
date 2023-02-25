package config

import (
	"github.com/spf13/viper"
)

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Options  map[string][]string
}

var db Database

func DB() *Database {
	return &db
}

func LoadDB() {
	db = Database{
		Name:     viper.GetString("database.name"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Options:  viper.GetStringMapStringSlice("database.options"),
	}
	// fmt.Println(db)
}
