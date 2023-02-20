package cmd

import (
	"PharmaProject/amqp"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"send"},
	// SuggestFor: []string{"start"},
	// PreRun: func(cmd *cobra.Command, args []string) {
	// },
	Short: "Initiates the Pharmacy Management System",
	Run: func(cmd *cobra.Command, args []string) {
		amqp.Send()
		// amqp.Receive()
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
