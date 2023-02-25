package config

import (
	"github.com/spf13/viper"
)

type AMQPConfig struct {
	URI string
}

var amqp AMQPConfig

// AMQP returns the default AMQP configuration
func AMQP() AMQPConfig {
	return amqp
}
func LoadAmqp() {
	amqp = AMQPConfig{
		URI: viper.GetString("amqp.uri"),
	}
}
