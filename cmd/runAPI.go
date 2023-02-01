package cmd

import (
	api "PharmaProject/api"

	"github.com/spf13/cobra"
)

var startAPICmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r", "ru"},
	// SuggestFor: []string{"start"},
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	api.Init()
	// },
	Short: "Initiates the Pharmacy Management System",
	Run: func(cmd *cobra.Command, args []string) {
		api.Init()
	},
}

func init() {
	rootCmd.AddCommand(startAPICmd)
}
