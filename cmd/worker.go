package cmd

import (
	"PharmaProject/worker"

	"github.com/spf13/cobra"
)

var startWorker = &cobra.Command{
	Use:     "worker",
	Aliases: []string{"w", "work"},
	// SuggestFor: []string{"start"},
	// PreRun: func(cmd *cobra.Command, args []string) {
	// },
	Short: "Initiates the Pharmacy Management System",
	Run: func(cmd *cobra.Command, args []string) {
		worker.ReceiveTask()
	},
}

func init() {
	rootCmd.AddCommand(startWorker)
}
