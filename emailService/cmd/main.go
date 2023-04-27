package main

import (
	"log"
	"mailService/broker"
	"mailService/config"
	"time"

	"github.com/memphisdev/memphis.go"
)

var (
	stationName  = "authStation"
	consumerName = "otpProducer"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Panic(err.Error())
	}

	// Connect to memphis broker
	conn, err := memphis.Connect(config.MEMPHIS_HOST, config.MEMPHIS_USERNAME, memphis.Password(config.MEMPHIS_PASSWORD))
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
	err = broker.ConsumeMessage(consumer, config)
	if err != nil {
		log.Println("3" + err.Error())
	}

}
