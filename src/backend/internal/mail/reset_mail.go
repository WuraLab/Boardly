package mail

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Minimum required to send an email
func resetEmail() []byte {
	address := "test@example.com"
	name := "Example User"
	from := mail.NewEmail(name, address)
	subject := "Hello World from the SendGrid Go Library"
	address = "test@example.com"
	name = "Example User"
	to := mail.NewEmail(name, address)
	content := mail.NewContent("text/plain", "some text here")
	m := mail.NewV3MailInit(from, subject, to, content)
	address = "test2@example.com"
	name = "Example User"
	email := mail.NewEmail(name, address)
	m.Personalizations[0].AddTos(email)
	return mail.GetRequestBody(m)
}

func SendResetEmail() {
	request := sendgrid.GetRequest(os.Getenv("YOUR_SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = resetEmail()
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
