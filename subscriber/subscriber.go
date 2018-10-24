package main

// Import Go and NATS packages
import (
	"log"
	"runtime"
	//"strconv"
	"github.com/nats-io/go-nats"
	//"encoding/json"
)

type Sensors struct {
	Name string
	Timestamp string
	Value string
 }

 type List struct {
	Sensor1 *Sensors
	Sensor2 *Sensors
	Sensor3 *Sensors
 }


func main() {
	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(natsConnection, nats.JSON_ENCODER)
	
	log.Println("Connected to " + nats.DefaultURL)
	//var sensor2 Sensors
	//log.Printf("Received message %+v\n", sensor2 )
	//var x int
	
	var list1 *List

	// Subscribe to subject
	log.Printf("Subscribing to subject 'foo'\n")
	c.Subscribe("foo", func(list *List) {
		//json.Unmarshal(msg.Data,&list.sensor1)
		// Handle the message
		log.Printf("Received message\n %+v\n %+v\n %+v\n", list.Sensor1, list.Sensor2, list.Sensor3 )
		//log.Printf("Received message %s %d %f\n", list.sensor2.name, list.sensor2.timestamp, list.sensor2.value )
		//log.Printf("Received message %s %d %f\n", list.sensor3.name, list.sensor3.timestamp, list.sensor3.value )
		list1 = list;

	})

	connectSensorDb()
	insertSensortable(list1)


	




	// Keep the connection alive
	runtime.Goexit()
}