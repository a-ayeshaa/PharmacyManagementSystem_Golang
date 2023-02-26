package conn

import "PharmaProject/internal/config"

var defaultAssignWorker Worker

// AssignWorker returns default worker
func DefaultAssignWorker() Worker {
	return defaultAssignWorker
}

// ConnectAssignWorker sets the client of worker using default configuration file
func ConnectAssignWorker() error {
	wCfg := config.Worker()
	wCfg.JobQueue += "_assign"
	wCfg.BindingKey += "_assign"
	aCfg := config.AMQP()
	rCfg := config.Redis()
	return defaultAssignWorker.Connect(&wCfg, &aCfg, &rCfg)
}