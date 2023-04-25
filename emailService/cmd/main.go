package main

import (
	"log"
	"mailService/broker"
	"mailService/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Panic(err.Error())
	}

	err = broker.ConsumeMessage(config)
	if err != nil {
		log.Panic(err.Error())
	}
}
