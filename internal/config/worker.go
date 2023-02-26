package config

import (
	"time"

	"github.com/spf13/viper"
)

// WorkerCfg holds the configuration of the worker
type WorkerCfg struct {
	JobQueue      string
	ResultTTL     time.Duration
	ExchangeName  string
	ExchangeType  string
	BindingKey    string
	PrefetchCount int
}

var worker WorkerCfg

// Worker returns the default Worker configuration
func Worker() WorkerCfg {
	return worker
}

// LoadWorker loads Worker configuration
func LoadWorker() {
	worker = WorkerCfg{
		JobQueue:      viper.GetString("worker.job_queue"),
		ResultTTL:     viper.GetDuration("worker.result_ttl") * time.Second,
		ExchangeName:  viper.GetString("worker.exchange_name"),
		ExchangeType:  viper.GetString("worker.exchange_type"),
		BindingKey:    viper.GetString("worker.binding_key"),
		PrefetchCount: viper.GetInt("worker.prefetch_count"),
	}
}
