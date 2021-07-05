package common

import (
	"crypto/tls"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail/v2"
)

func SendMail(subject string, senderEmail string, htmlBody string) {
	server := mail.NewSMTPClient()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// SMTP Server
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = os.Getenv("EMAIL_USER")
	server.Password = os.Getenv("EMAIL_PASSWORD")
	server.Encryption = mail.EncryptionSTARTTLS
	SetEmailFrom := os.Getenv("SetEmailFrom")

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
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.SetFrom(SetEmailFrom).AddTo(senderEmail).SetSubject(htmlBody)

	email.SetBody(mail.TextHTML, htmlBody)

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}