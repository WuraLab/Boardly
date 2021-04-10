package main

import (
	"os"
	"log"
	"github.com/xhit/go-simple-mail/v2"
)

func sendEmail(sender_email string, subject string, html_body string) {
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = os.Getenv("EMAIL_HOST")
	server.Password = os.Getenv("EMAIL_PASSWORD")
	server.Encryption = mail.EncryptionSTARTTLS

	// Since v2.3.0 you can specified authentication type:
	// - PLAIN (default)
	// - LOGIN
	// - CRAM-MD5
	// server.Authentication = mail.AuthPlain

	// Variable to keep alive connection
	server.KeepAlive = true

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient,err := server.Connect()

	if err != nil{
		log.Fatal(err)
	}


	email := mail.NewMSG()
	email.SetFrom("Test").
		AddTo(sender_email).
		SetSubject(subject)

	email.SetBody(mail.TextHTML, html_body)

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}