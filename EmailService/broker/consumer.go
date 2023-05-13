package broker

import (
	"context"
	"fmt"
	"mailService/config"
	"mailService/mailer"

	"github.com/memphisdev/memphis.go"
)

func ConsumeMessage(consumer *memphis.Consumer, config *config.Config) error {

	handler := func(msgs []*memphis.Msg, err error, ctx context.Context) {
		for _, msg := range msgs {
			// fmt.Println(string(msg.Data()))
			msg.Ack()
			headers := msg.GetHeaders()
			// fmt.Println(headers["email"])
			// fmt.Println(headers["otp"])

			err = mailer.SendMail(headers["email"], "EventHub OTP", headers["otp"], *config)
			if err != nil {
				fmt.Printf("Send mail failed: %s\n", err.Error())
				continue
			}
		}
	}

	forever := make(chan bool)

	err := consumer.Consume(handler)
	if err != nil {
		return err
	}

	// Reference : https://stackoverflow.com/questions/47262088/golang-forever-channel
	<-forever // To run the consumer forever

	return nil
}
