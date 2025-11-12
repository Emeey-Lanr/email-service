package message_broker

import (
	"fmt"
	"time"

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

			// since sendEmail return a bool
			// when it's  true, we break the loop
			// and change success to true if we break
			//  if unable to break, success stays false 
			// and it's used to either says it successful or publish to the dead queue

		    for i := 0; i <= maxRetry; i++ {
            //    err := Send_email() // err is either true or false
			//    if !err {
			// 	*successAddress = true
            //       break // exit loop
			//    }

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
