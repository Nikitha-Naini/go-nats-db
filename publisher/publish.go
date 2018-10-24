package main

import (
	"log"
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
	"math/rand"
	"strconv"
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

	var list *List

	var s1 *Sensors
	var s2 *Sensors
	var s3 *Sensors
	now := time.Now().Unix()
	val := rand.Int63n(100)
	s1 = &Sensors{Name: "Sensor1", Timestamp: strconv.FormatInt(now,10), Value: strconv.FormatInt(val, 10)}
	val = rand.Int63n(100)
	s2 = &Sensors{Name: "Sensor1", Timestamp: strconv.FormatInt(now,10), Value: strconv.FormatInt(val, 10)}
	val = rand.Int63n(100)
	s3 = &Sensors{Name: "Sensor1", Timestamp: strconv.FormatInt(now,10), Value: strconv.FormatInt(val, 10)}

	list = &List{Sensor1: s1, Sensor2: s2, Sensor3: s3}

	//m, err := json.Marshal(list.sensor1)

	//fmt.Println(err)
	log.Printf("Sent message %+v %+v %+v\n", list.Sensor1,list.Sensor2, list.Sensor3 )
	fmt.Println("Publishing Hello World")

	// publisher connecting to nats server 

	natsConnection, err := nats.Connect(nats.DefaultURL)
	c, err1 := nats.NewEncodedConn(natsConnection, nats.JSON_ENCODER)

	if(err != nil) {
	fmt.Println(err)
	}
	if(err1 != nil) {
	fmt.Println(err1)
	}
	defer c.Close()
	fmt.Println("Connected to NATS server: " + nats.DefaultURL)
	//var x = 10
	// Msg structure
	/*msg := &nats.Msg{
		Subject: "foo",
		Reply:   "bar",
		Data:    m,
	}*/

	c.Publish("foo",list)
	//c.Publish("foo",list.sensor2)
	//c.Publish("foo",list.sensor3)

	//var sensorx Sensors
	//	var list1 List
	//err1 := json.Unmarshal(msg.Data,&sensorx)
	//fmt.Println(err1)

	//log.Println("Published msg.Subject = foo | msg.Data = ")
	//log.Printf("Published message %+v\n", list)
		//log.Printf("Published message %s %d %f\n", list1.sensor2.name, list1.sensor2.timestamp, list1.sensor2.value )
		//log.Printf("Published message %s %d %f\n", list1.sensor3.name, list1.sensor3.timestamp, list1.sensor3.value )

		err = natsConnection.Flush()
	if(err == nil) {
		// Everything has been processed by the server for nc *Conn.
	} else {
		fmt.Println(err)
	}

	err = c.Flush()
	if(err == nil) {
		// Everything has been processed by the server for nc *Conn.
	} else {
		fmt.Println(err)
	}
}