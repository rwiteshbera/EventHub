package main

import (
	"log"
	"mailService/broker"
	"mailService/config"
	"time"

	"github.com/memphisdev/memphis.go"
)

var (
	stationName  = "auth"
	consumerName = "otpConsumer"
)

func main() {
	Config := config.LoadConfig()

	// Connect to memphis broker
	conn, err := memphis.Connect(Config.MEMPHIS_HOST, Config.MEMPHIS_USERNAME, memphis.Password(Config.MEMPHIS_PASSWORD))
	if err != nil {
		log.Println("1" + err.Error())
	}
	defer conn.Close()

	// Create a new consumer
	consumer, err := conn.CreateConsumer(stationName, consumerName, memphis.PullInterval(1*time.Second))
	if err != nil {
		log.Println("2" + err.Error())
	}

	// Consume messages and send email
	err = broker.ConsumeMessage(consumer, Config)
	if err != nil {
		log.Println("3" + err.Error())
	}

}
