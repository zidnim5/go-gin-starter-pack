package config

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection

func GetBrokerConn() *amqp.Connection {
	if conn == nil {
		conn_tmp, err := amqp.Dial("amqp://guest:guest@localhost:" + os.Getenv("MESSAGE_BROKER_PORT") + "/")
		if err != nil {
			fmt.Println("Failed Initializing Broker Connection")
			panic(err)
		}
		conn = conn_tmp
	}
	return conn
}
