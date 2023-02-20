package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"PharmaProject/connection"
	"PharmaProject/helper"
	"context"
	"log"
	"time"
)

func Send() {
	conn:=connection.AMQPCon()
	defer conn.Close()
	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"complains", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	  )
	  helper.FailOnError(err, "Failed to declare a queue")
	  
	  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	  defer cancel()
	  
	  body := "Hello World!"
	  err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "application/plain",
		//   ContentEncoding: "encoding/csv",
		  Body:        []byte(body),
		})
		helper.FailOnError(err, "Failed to publish a message")
	  log.Printf(" [x] Sent %s\n", body)
}
