package main

import (
	"github.com/joho/godotenv"
	"github.com/emeey-lanr/email_service/message_broker"
	"log"
)


func main (){
	
   if  err := godotenv.Load(); err != nil{
	log.Fatal("Unable to connect tp dotenv")
   }
  
// Connect to RabbitMQ
   connection, channel := message_broker.Connect_to_rabitmq()

   defer connection.Close()
   defer channel.Close()
    
  
//  Consume messages from queue


}