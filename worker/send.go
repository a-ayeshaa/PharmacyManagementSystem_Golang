package amqp

import (
	"PharmaProject/connection"
	"PharmaProject/helper"
	"context"
	"io"
	"io/ioutil"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type QueueConfig struct {
	name       string
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
}

func SendTask(meds io.ReadCloser) {
	conn := connection.AMQPCon()
	defer conn.Close()
	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")

	defer ch.Close()
	q, err := ch.QueueDeclare(
		viper.GetString("worker.name"),
		viper.GetBool("worker.durable"),
		viper.GetBool("worker.autoDelete"),
		viper.GetBool("worker.exclusive"),
		viper.GetBool("worker.noWait"),
		nil,
	)
	helper.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body,err := ioutil.ReadAll(meds)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	helper.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
