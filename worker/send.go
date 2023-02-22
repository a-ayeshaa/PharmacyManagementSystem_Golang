package worker

import (
	"PharmaProject/connection"
	"PharmaProject/helper"
	"context"
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

func SendTask(meds []byte) {
	// fmt.Println(meds)
	conn := connection.AMQPCon()
	defer conn.Close()
	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")

	defer ch.Close()
	// Direct exchange ....
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

	// body, err := ioutil.ReadAll(meds)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        meds,
		})
	helper.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", meds)

	// Fanout exchange ...
	// err = ch.ExchangeDeclare(
	// 	"logs",   // name
	// 	"fanout", // type
	// 	true,     // durable
	// 	false,    // auto-deleted
	// 	false,    // internal
	// 	false,    // no-wait
	// 	nil,      // arguments
	// )
	// helper.FailOnError(err, "Failed to declare an exchange")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// // body, err := ioutil.ReadAll(meds)
	// err = ch.PublishWithContext(ctx,
	// 	"logs", // exchange
	// 	"",     // routing key
	// 	false,  // mandatory
	// 	false,  // immediate
	// 	amqp.Publishing{
	// 		ContentType: "application/json",
	// 		Body:        meds,
	// 	})
	// helper.FailOnError(err, "Failed to publish a message")

	// log.Printf(" [x] Sent %s", meds)
}
