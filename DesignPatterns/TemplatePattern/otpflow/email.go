package otpflow

import "fmt"

type Email struct {
}

func (e *Email) generateRandomOtp() string {
	return "12345"
}

func (e *Email) saveOtpInCache(otp string) {
	fmt.Printf("Saving OTP in cache: %s\n", otp)
}

func (e *Email) getMessage(otp string) string {
	msg := fmt.Sprintf("Your OTP is %s and it will expire in 10 minutes \n", otp)
	return msg
}

func (e *Email) sendNotification(msg string) error {
	fmt.Printf("Email sent : %s\n", msg)
	return nil
}
