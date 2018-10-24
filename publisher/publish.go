package main

import (
	"log"
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
	"math/rand"
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

	var list List

	now := time.Now().Unix()
	value := rand.Float64()
	list.sensor1 = Sensors{"Sensor1",now,value}
	value = rand.Float64()
	list.sensor2 = Sensors{"Sensor2",now,value}
	value = rand.Float64()
	list.sensor3 = Sensors{"Sensor3",now,value}

	m, err := json.Marshal(list.sensor1)

	fmt.Println(err)
	log.Printf("Sent message %s %d %f\n", list.sensor1.name, list.sensor1.timestamp, list.sensor1.value )
	fmt.Println("Publishing Hello World")

	// publisher connecting to nats server 

	natsConnection, _ := nats.Connect(nats.DefaultURL)
	defer natsConnection.Close()
	fmt.Println("Connected to NATS server: " + nats.DefaultURL)

	// Msg structure
	msg := &nats.Msg{
		Subject: "foo",
		Reply:   "bar",
		Data:    m,
	}
	natsConnection.PublishMsg(msg)

	var sensorx Sensors

	err1 := json.Unmarshal(msg.Data,&sensorx)
	fmt.Println(err1)

	log.Println("Published msg.Subject = "+msg.Subject, "| msg.Data = ")
	log.Printf("Received message %s %d %f\n", sensorx.name, sensorx.timestamp, sensorx.value )
		//log.Printf("Received message %s %d %f\n", list1.sensor2.name, list1.sensor2.timestamp, list1.sensor2.value )
		//log.Printf("Received message %s %d %f\n", list1.sensor3.name, list1.sensor3.timestamp, list1.sensor3.value )


}