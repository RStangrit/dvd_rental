package emailclient

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmail(sender, receiver, subject, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("mailhog", 1025, "", "")

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Failed to send email: %v", err)
		return
	}

	fmt.Println("Email sent!")
}
