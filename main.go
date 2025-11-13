package main

import (
	"github.com/joho/godotenv"
	"github.com/emeey-lanr/email_service/message_broker"
	"log"
   "net/http"
)


func main (){
	
   if  err := godotenv.Load(); err != nil{
	log.Println("Unable to connect to dotenv")
   }

   go func ()  {
         http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("Email service is running"))
   })


    log.Println("Email running on port 8080")
     if err := http.ListenAndServe(":8080", nil); err != nil{
      log.Fatal(err)
     }
   
   }()

   
  
// Connect to RabbitMQ
   connection, channel := message_broker.Connect_to_rabitmq()

   defer connection.Close()
   defer channel.Close()

// Consumer
    message_broker.Consumer(channel)

    

}