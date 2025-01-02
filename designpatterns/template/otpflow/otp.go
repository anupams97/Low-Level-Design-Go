package otpflow

type IOtp interface {
	generateRandomOtp() string
	saveOtpInCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type OTP struct {
	// any one that implement IOtp can be here
	IOtp IOtp
}

func (o *OTP) GenAndSendOtp() error {
	otp := o.IOtp.generateRandomOtp()
	o.IOtp.saveOtpInCache(otp)
	msg := o.IOtp.getMessage(otp)
	return o.IOtp.sendNotification(msg)
}
