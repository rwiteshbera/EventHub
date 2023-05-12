package broker

import (
	"authenticationService/config"
	"fmt"

	"github.com/memphisdev/memphis.go"
)

// Send OTP to Memphis
func ProduceMessage(email, otp string, config *config.Config) bool {
	conn, err := memphis.Connect(config.MEMPHIS_HOST, config.MEMPHIS_USERNAME, memphis.Password(config.MEMPHIS_PASSWORD))
	if err != nil {
		return false
	}
	defer conn.Close()
	p, err := conn.CreateProducer("auth", "otpProducer")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	hdrs := memphis.Headers{}
	hdrs.New()
	err = hdrs.Add("email", email)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	err = hdrs.Add("otp", otp)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	err = p.Produce([]byte("You have a message!"), memphis.MsgHeaders(hdrs))

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
