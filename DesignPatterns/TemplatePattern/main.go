package main

import (
	"awesomeProject/TemplatePattern/otpflow"
	"fmt"
)

func main() {

	sms := otpflow.Sms{}
	email := otpflow.Email{}

	otp := otpflow.OTP{
		IOtp: &sms,
	}
	err := otp.GenAndSendOtp()
	if err != nil {
		return
	}

	otp = otpflow.OTP{
		IOtp: &email,
	}
	err = otp.GenAndSendOtp()
	if err != nil {
		return
	}

	smsOtp := otpflow.OTP{IOtp: &otpflow.Sms{}}
	emailOtp := otpflow.OTP{IOtp: &otpflow.Email{}}

	if err := smsOtp.GenAndSendOtp(); err != nil {
		fmt.Println("Failed to send SMS OTP:", err)
	}

	if err := emailOtp.GenAndSendOtp(); err != nil {
		fmt.Println("Failed to send Email OTP:", err)
	}

}
