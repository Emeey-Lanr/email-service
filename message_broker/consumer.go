package message_broker

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/emeey-lanr/email_service/email"
	"github.com/emeey-lanr/email_service/model"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)


func Consumer (channel *amqp.Channel, rds *redis.Client) {

//   Email Queue
	_, err := channel.QueueDeclare(
		 "email.queue",
		 true,
		 false,
		 false,
		 false,
		 nil,
	)


	 RabbitMqError(err, "failed to declare email queue")

	//  Dead-Letter (Failed Queue)	
  _, err = channel.QueueDeclare(
	"failed.queue",
	true,
	false,
	false,
	false,
  nil,
  )

  	 RabbitMqError(err, "Failed to declare failed queue")


	 messages, err := channel.Consume(
		"email.queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	 )
 	
	 
	 RabbitMqError(err, "Failed to register consumer")

	forever := make(chan bool)
 
	maxRetry := 3

	go func (){
		for data := range messages {
			success := false
            successAddress := &success
           
			var messagequed model.QueueResponse
        
			json.Unmarshal([]byte(data.Body), &messagequed)

			// since sendEmail return a bool
			// when it's  true, we break the loop
			// and change success to true if we break
			//  if unable to break, success stays false 
			// and it's used to either says it successful or publish to the dead queue

		    for i := 0; i <= maxRetry; i++ {
              err := email.Send_Email(messagequed.Email, messagequed.Data.Variable.Name, messagequed.Data.Template_code, rds) // err is either true or false
			   if !err {
				*successAddress = true
                  break // exit loop
			   }

			    *successAddress = false
				time.Sleep(time.Second * 60) // waits for 60 seconds before retrying 
			    

             if success {
				 fmt.Println("message sent succesfully")
				data.Ack(false) // accknowledge message sent succesfully
				// make a post request to change status
			 }else{
				
				fmt.Println("Failed to send message")
				//publish to failed queue
			 }

			}
		// your sen
          fmt.Println("received message", data.Body)
		   
		  
		}
	
	}()

	<- forever

}
