package db

import (
	config "PharmaProject/config"

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
	// viper.SetConfigName("config")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	config.Init()
	db = Database{
		Name:     viper.GetString("database.name"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Options:  viper.GetStringMapStringSlice("database.options"),
	}
}

//If .env is used for database configuration..

// type Config struct {
// 	DBSource   string `mapstructure:"DB_SOURCE"`
// 	DBUser     string `mapstructure:"DB_USER"`
// 	DBPassword string `mapstructure:"DB_PASSWORD"`
// 	DBName     string `mapstructure:"DB_NAME"`
// 	DBPort     string `mapstructure:"DB_PORT"`
// 	DBHost     string `mapstructure:"DB_HOST"`
// }

//

//

// func LoadConfig() (*Config, error) {
// 	viper.SetConfigFile(".env")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var config Config
// 	err = viper.Unmarshal(&config)
// 	return &config, nil
// }
