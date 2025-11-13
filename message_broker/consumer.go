package message_broker

import (
	"encoding/json"
	"log"
	"time"
	"github.com/emeey-lanr/email_service/email"
	"github.com/emeey-lanr/email_service/model"
	amqp "github.com/rabbitmq/amqp091-go"

)


func Consumer (channel *amqp.Channel) {

//   Email Queue
	_, err := channel.QueueDeclare(
		 "email.queue",
		 true,
		 false,
		 false,
		 false,
		 nil,
	)


	log.Println("failed to declare email.queu")

	//  Dead-Letter (Failed Queue)	
  _, err = channel.QueueDeclare(
	"failed.queue",
	true,
	false,
	false,
	false,
  nil,
  )

  	log.Println("Failed to declare failed queue")


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
            log.Println(string(data.Body), "message body")

			var messagequed model.QueueResponse
        
			 err:= json.Unmarshal([]byte(data.Body), &messagequed)
			 if err != nil {
				log.Println("Error decoding json", err)
				continue
			 }
			
             
			//retry 3 times
		    for i := 0; i <= maxRetry; i++ {
				
			// we send a bool true if there's error and false if there's no error
              err := email.Send_Email(messagequed) // err is either true or false
			   if !err  {
				*successAddress = true
                  break // exit loop
			   }

			    *successAddress = false
				time.Sleep(time.Second * 60) // waits for 60 seconds before retrying 
			    

             if success {
				 log.Println("message sent succesfully")
				data.Ack(false) // accknowledge message sent succesfully
				// make a post request to change status
			 }else{
				channel.Publish(
					"",
					"failed.queue",
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body: data.Body,
					},
				)
				log.Println("Failed to send message, published to failed queue")
				//publish to failed queue
			 }

			}
		  
		}
	
	}()

	<- forever

}
