package main

import "fmt"

// start template method
type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

// defines base template algorithm
func (o *Otp) genAndSendOtp(otpLenght int) error {
	otp := o.iOtp.genRandomOTP(otpLenght)
	o.iOtp.saveOTPCache(otp)
	msg := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(msg)
	if err != nil {
		return err
	}
	return nil
}

// end template method

// concrete implementation
type sms struct{}

func (s *sms) genRandomOTP(len int) string {
	randomOtp := "123123"
	fmt.Printf("SMS: generating random otp %s\n", randomOtp)
	return randomOtp
}

func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(msg string) error {
	fmt.Printf("SMS: sending sms: %s\n", msg)
	return nil
}

// another concrete implementation
type email struct{}

func (s *email) genRandomOTP(len int) string {
	randomOtp := "9292929"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOtp)
	return randomOtp
}

func (s *email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *email) sendNotification(msg string) error {
	fmt.Printf("EMAIL: sending email: %s\n", msg)
	return nil
}

func main() {
	otp := &Otp{
		iOtp: &sms{},
	}
	otp.genAndSendOtp(4)

	fmt.Println()

	otp.iOtp = &email{}
	otp.genAndSendOtp(5)
}
