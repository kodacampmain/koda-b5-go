package users

import "fmt"

// Service that sends emails
type EmailService struct{}

func (es *EmailService) Send(to, message string) {
	fmt.Printf("Sending email to %s: %s\n", to, message)
}

// Service that sends WhatsApp
type WhatsappService struct{}

func (ws *WhatsappService) Send(to, message string) {
	fmt.Printf("Sending whatsapp message to %s: %s\n", to, message)
}

// Service that sends SMS
type SMSService struct{}

func (ss *SMSService) Send(to, message string) {
	fmt.Printf("Sending SMS message to %s: %s\n", to, message)
}

type IMessageService interface {
	Send(to, message string)
}

// UserService that depends directly on EmailService
type UserService struct{}

func (us *UserService) RegisterUser(msgService IMessageService, to, message string) {
	// smsService := &SMSService{}
	msgService.Send(to, message)
}

func main() {
	var email string = "alice@example.com"
	var phone string = "081789012345"

	userService := &UserService{}

	emailService := &EmailService{}
	userService.RegisterUser(emailService, email, "Welcome to our app")

	// waService := &WhatsappService{}
	// userService.RegisterUser(waService, phone, "Welcome to our app")

	smsService := &SMSService{}
	userService.RegisterUser(smsService, phone, "Welcome to our app")
}
