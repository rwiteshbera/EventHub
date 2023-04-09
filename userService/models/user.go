package models

type UserLogin struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OTP struct {
	Otp string `json:"otp"`
}
