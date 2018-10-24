package main

// Import Go and NATS packages
import (
	"log"
	"runtime"
	
	"github.com/nats-io/go-nats"
	"encoding/json"
)

type Sensors struct {
	name string
	timestamp int64
	value float64
 }

 type List struct {
	 sensor1 Sensors
	 sensor2 Sensors
	 sensor3 Sensors
 }


func main() {
	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	var list List

	// Subscribe to subject
	log.Printf("Subscribing to subject 'foo'\n")
	natsConnection.Subscribe("foo", func(msg *nats.Msg) {
		json.Unmarshal(msg.Data,&list.sensor1)
		// Handle the message
		log.Printf("Received message %s %d %f\n", list.sensor1.name, list.sensor1.timestamp, list.sensor1.value )
		//log.Printf("Received message %s %d %f\n", list.sensor2.name, list.sensor2.timestamp, list.sensor2.value )
		//log.Printf("Received message %s %d %f\n", list.sensor3.name, list.sensor3.timestamp, list.sensor3.value )
	})

	// Keep the connection alive
	runtime.Goexit()
}