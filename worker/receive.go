package amqp

import (
	"PharmaProject/connection"
	"PharmaProject/helper"
	"PharmaProject/models"
	con "PharmaProject/usecase"
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ReceiveTask() {
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	helper.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var meds []models.Medicine
			err := json.NewDecoder(bytes.NewReader(d.Body)).Decode(&meds)
			newmed, err := con.NewMedicine().AddBulkMedicine(meds)
			if err != nil {
				
			}
			fmt.Println(newmed)
		}

	}()

	//   log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
