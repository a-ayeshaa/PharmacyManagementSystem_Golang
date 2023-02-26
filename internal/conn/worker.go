package conn

import (
	"PharmaProject/config"
	"time"

	machinery "github.com/RichardKnop/machinery/v1"
	mchnrycfg "github.com/RichardKnop/machinery/v1/config"
)

// Worker holds the worker instace
type Worker struct{ *machinery.Server }

// defaultConfig is the default worker instance
var defaultWorker Worker

// Connect sets the db client of worker using configuration cfg
func (w *Worker) Connect(wCfg *config.WorkerCfg, aCfg *config.AMQPConfig, rCfg *config.RedisCfg) error {
	srv, err := machinery.NewServer(&mchnrycfg.Config{
		Broker:          aCfg.URI,
		DefaultQueue:    wCfg.JobQueue,
		ResultBackend:   rCfg.URI(),
		ResultsExpireIn: int(wCfg.ResultTTL / time.Second),
		AMQP: &mchnrycfg.AMQPConfig{
			Exchange:      wCfg.ExchangeName,
			ExchangeType:  wCfg.ExchangeType,
			BindingKey:    wCfg.BindingKey,
			PrefetchCount: wCfg.PrefetchCount,
		},
	})
	if err != nil {
		return err
	}
	w.Server = srv
	return nil
}

// DefaultWorker returns default worker
func DefaultWorker() Worker {
	return defaultWorker
}

// ConnectWorker sets the client of worker using default configuration file
func ConnectWorker() error {
	wCfg := config.Worker()
	aCfg := config.AMQP()
	rCfg := config.Redis()
	return defaultWorker.Connect(&wCfg, &aCfg, &rCfg)
}
