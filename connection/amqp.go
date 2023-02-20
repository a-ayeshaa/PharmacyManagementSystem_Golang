package connection

import (
	"PharmaProject/helper"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type AMQPConfig struct {
	URI string
}
func AMQPCon() *amqp.Connection {
	conn, err := amqp.Dial(viper.GetString("amqp.uri"))
	helper.FailOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	return conn
}
