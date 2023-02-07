package cmd

import (
	api "PharmaProject/api"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	model "PharmaProject/models"
)

var startAPICmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r", "ru"},
	// SuggestFor: []string{"start"},
	PreRun: func(cmd *cobra.Command, args []string) {
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  "host=localhost user=ayesha password=password dbname=pharmacy_db port=8080 sslmode=disable",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
		if err!=nil{
			panic(err)
		}
		db.AutoMigrate(model.User{},model.Cart{},model.Medicine{},model.Order{})
	},
	Short: "Initiates the Pharmacy Management System",
	Run: func(cmd *cobra.Command, args []string) {
		api.Init()
	},
}

func init() {
	rootCmd.AddCommand(startAPICmd)
}
