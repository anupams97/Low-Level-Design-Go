package otpflow

import "fmt"

type Sms struct {
}

func (s *Sms) generateRandomOtp() string {
	return "12345"
}

func (s *Sms) saveOtpInCache(otp string) {
	fmt.Printf("Saving OTP in cache: %s\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	msg := fmt.Sprintf("Your OTP is %s and it will expire in 10 minutes \n", otp)
	return msg
}

func (s *Sms) sendNotification(msg string) error {
	fmt.Printf("message sent : %s\n", msg)
	return nil
}
