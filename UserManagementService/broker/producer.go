package broker

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"userService/config"
)

type Body struct {
	Email string
	OTP   string
}

// Send OTP to Memphis
func ProduceMessage(email, otp string, config *config.Config) bool {
	conn, err := amqp091.Dial(config.RABBITMQ)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer ch.Close()

	// Create a queue to send message
	queue, err := ch.QueueDeclare(
		"otp-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	err = ch.PublishWithContext(
		context.Background(),
		"",
		queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			MessageId:   email,
			Body:        []byte(otp),
		})
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
