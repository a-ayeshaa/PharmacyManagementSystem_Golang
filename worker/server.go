package worker

import (
	"PharmaProject/internal/config"
	"log"
	"time"

	machinery "github.com/RichardKnop/machinery/v1"
	mchnrycfg "github.com/RichardKnop/machinery/v1/config"
)

// Server returns default machinery
func Server() (*machinery.Server, error) {
	wCfg := config.Worker()
	aCfg := config.AMQP()
	rCfg := config.Redis()

	log.Println(aCfg.URI)

	srvr, err := machinery.NewServer(&mchnrycfg.Config{
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
		return nil, err
	}

	if err := srvr.RegisterTasks(add_med_tasks); err != nil {
		return nil, err
	}
	return srvr, nil
}
