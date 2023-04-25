package broker

import (
	"context"
	"mailService/config"
	"mailService/mailer"
	"time"

	"github.com/memphisdev/memphis.go"
)

func ConsumeMessage(config *config.Config) error {
	conn, err := memphis.Connect(config.MEMPHIS_HOST, config.MEMPHIS_USERNAME, memphis.Password(config.MEMPHIS_PASSWORD))
	if err != nil {
		return err
	}
	defer conn.Close()

	consumer, err := conn.CreateConsumer("authStation", "otpProducer", memphis.PullInterval(1*time.Second))
	if err != nil {
		return err
	}

	handler := func(msgs []*memphis.Msg, err error, ctx context.Context) {
		for _, msg := range msgs {
			// fmt.Println(string(msg.Data()))
			msg.Ack()
			headers := msg.GetHeaders()
			// fmt.Println(headers["email"])
			// fmt.Println(headers["otp"])

			mailer.SendMail(headers["email"], "EventHub OTP", headers["otp"], *config)
		}
	}

	err = consumer.Consume(handler)
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Minute)
	return nil
}
