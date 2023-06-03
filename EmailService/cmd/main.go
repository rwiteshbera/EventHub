package main

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"mailService/config"
	"mailService/mailer"
)

func main() {
	Config := config.LoadConfig()

	conn, err := amqp091.Dial(Config.RABBITMQ)
	fmt.Println(Config.RABBITMQ)
	if err != nil {
		log.Println(err.Error())
	}
	defer func(conn *amqp091.Connection) {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err.Error())
	}
	defer func(ch *amqp091.Channel) {
		err := ch.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(ch)

	messages, err := ch.Consume(
		"otp-queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err.Error())
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			err := mailer.SendMail(d.MessageId, "EventHub OTP Verification", string(d.Body), Config)
			if err != nil {
				return
			}
		}
	}()
	<-forever
}
