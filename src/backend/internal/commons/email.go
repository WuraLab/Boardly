package commons

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/wuraLab/boardly/src/backend/internal/config"
	mail "github.com/xhit/go-simple-mail/v2"
)

func SendMail(subject string, senderEmail string, htmlBody string) bool {

	var Config config.Mail
	var err error

	//Load configuration
	// if Config, err = config.LoadConfig(".env"); err != nil {
	// 	log.Fatal(err)
	// }
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = Config.Email
	server.Password = Config.Password
	server.Encryption = mail.EncryptionSTARTTLS

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
	email.SetFrom(Config.SetEmailFrom).
		AddTo(senderEmail).
		SetSubject(subject)

	email.SetBody(mail.TextHTML, htmlBody)

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
		return false
	} else {
		log.Println("Email Sent")
		return true
	}
}
