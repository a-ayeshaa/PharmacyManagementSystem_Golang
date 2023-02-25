package cmd

import (
	"PharmaProject/conn"
	"PharmaProject/worker"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:     "worker",
	Aliases: []string{"w", "work"},
	// SuggestFor: []string{"start"},
	PreRun: func(cmd *cobra.Command, args []string) {
		// config.init()

		// need to connect worker for requeue purpose
		if err := conn.ConnectWorker(); err != nil {
			panic(fmt.Errorf("Can't connect worker: %v", err))
		}

		// need to connect assign worker for requeue purpose
		if err := conn.ConnectAssignWorker(); err != nil {
			panic(fmt.Errorf("Can't connect assign worker: %v", err))
		}

		// connect redis
		if err := conn.ConnectRedis(); err != nil {
			panic(fmt.Errorf("Can't connect redis: %v", err))
		}
	},
	Short: "Initiates the worker",
	Run: func(cmd *cobra.Command, args []string) {
		startWorker()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}

func startWorker() error{
	srvr, err := worker.Server()
	if err != nil {
		return err
	}
	if err := srvr.NewWorker("machinery_worker", 1).Launch(); err != nil {
		log.Println(err)
	}

	return nil
}
