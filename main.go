package main

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func main() {
	sender := ""
	password := ""
	smtpHost := "smtp.google.com"
	smtpPort := 587

	to := ""
	subject := "Hello"

	mail := gomail.NewMessage()
	mail.SetHeader("From", sender)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", "<h1>Hai</h1>")

	dialer := gomail.NewDialer(smtpHost, smtpPort, sender, password)

	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Println("Gagal mengirim email : ", err)
		os.Exit(1)
	}

}
