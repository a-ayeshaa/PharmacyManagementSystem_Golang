package cmd

import (
	pharma "PharmaProject/pkg"
    "github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
    Use:   "start",
	Aliases: []string{"st","s"},
	// SuggestFor: []string{"start"},
    Short:  "Initiates the Pharmacy Management System",
    Run: func(cmd *cobra.Command, args []string) {
        pharma.Start()
    },
}

func init() {
    rootCmd.AddCommand(startCmd)
}
