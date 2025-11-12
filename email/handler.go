package email

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/redis/go-redis/v9"

)


func Send_Email (email, name, template_code string, rds *redis.Client) bool{
 
   
    sbj, bdy := Decode_template(name, template_code,  rds)

	from := mail.NewEmail("App", "oyelowo.emmauel001@gmail.com")
	subject := sbj
	to := mail.NewEmail(name, email)
	plainTextContent := bdy
	htmlContent := fmt.Sprintf("<h1>%s</h1>", bdy)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDMAILGRID_API_KEY"))

	response, err := client.Send(message)

	// if there is an error it returns true
	// if there is no error it returns fasle
	if err != nil{
 
		return true
	}

	fmt.Println(response)

	return false
} 