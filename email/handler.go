package email

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
		"github.com/emeey-lanr/email_service/model"

)


func Send_Email (email model.QueueResponse) bool{
 

	from := mail.NewEmail("App", "oyelowo.emmanuel001@gmail.com")
	subject := email.Subject
	to := mail.NewEmail(email.Data.Variable.Name, email.Email)
	plainTextContent := email.Text_body
	htmlContent := email.Html_body

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDMAILGRID_API_KEY"))

	response, err := client.Send(message)

	// if there is an error it returns true
	// if there is no error it returns fasle
	if err != nil{
		log.Println("Unable to send email", err)
		return true
	}

	log.Println(response)

	return false
} 