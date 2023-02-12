package cmd

import (
	api "PharmaProject/api"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var startAPICmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s", "ser","serv"},
	// SuggestFor: []string{"start"},
	// PreRun: func(cmd *cobra.Command, args []string) {
	// },
	Short: "Initiates the Pharmacy Management System",
	Run: func(cmd *cobra.Command, args []string) {
		r := api.Init()
		port := ":3000"
		fmt.Println("Listening to port ", port)
		http.ListenAndServe(port, r)
	},
}

func init() {
	rootCmd.AddCommand(startAPICmd)
}
