package cmd

import (
	api "PharmaProject/api"
	"PharmaProject/conn"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var startAPICmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s", "ser", "serv"},
	// SuggestFor: []string{"start"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// config.init()

	},
	Short: "Initiates the Pharmacy Management System",
	Run: func(cmd *cobra.Command, args []string) {
		r := api.Init()
		if err := conn.ConnectWorker(); err != nil {
			fmt.Printf("failed to connect worker: %v", err)
		}
		if err := conn.ConnectRedis(); err != nil {
			panic(fmt.Errorf("Can't connect redis: %v", err))
		}
		port := ":3000"
		fmt.Println("Listening to port ", port)
		http.ListenAndServe(port, r)
	},
}

func init() {
	rootCmd.AddCommand(startAPICmd)
}
