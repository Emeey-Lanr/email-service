package main

import (
	"github.com/joho/godotenv"
	"github.com/emeey-lanr/email_service/message_broker"
   "github.com/emeey-lanr/email_service/cache"
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

   rds := cache.ConnectToRedis()
    
       message_broker.Consumer(channel, rds)
//  Consume messages from queue
    

}