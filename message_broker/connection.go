package message_broker

import (
	"os"
	amqp "github.com/rabbitmq/amqp091-go"
	
)

func Connect_to_rabitmq () (*amqp.Connection, *amqp.Channel){

	Url := os.Getenv("RABBITMQ_URL")
	connection, err := amqp.Dial(Url)
   RabbitMqError(err, "failed to open connection to Rbbitmq")

    channel , err := connection.Channel()
	  RabbitMqError(err, "failed to connect to channel")




 return connection, channel
   

}