package main

import (
  "database/sql"
  "fmt"
	"strconv"
  _ "github.com/lib/pq"
)


const (
  host     = "localhost"
  port     = 5432
  user     = "naini"
  password = "password"
  dbname   = "sensor_db"
)

type tools struct {
    db *sql.DB
}
var t tools

func average(list *List) float64 {
	v1, _ := strconv.ParseInt(list.Sensor1.Value, 10, 64)
	v2, _ := strconv.ParseInt(list.Sensor2.Value, 10, 64)
	v3, _ := strconv.ParseInt(list.Sensor3.Value, 10, 64)
	avg := float64((v1 + v2 + v3)/3)
	return avg
}

func insertSensortable(list *List){
	
	avg := average(list)

	sqlStatement := fmt.Sprintf("INSERT INTO sensors (name, timestamp, value) VALUES (%+v, %+v, %+v)",list.Sensor1.Name, list.Sensor1.Timestamp, list.Sensor1.Value )
	db := t.db
	_, err := db.Exec(sqlStatement)
	if err != nil {
	  panic(err)
	}

	sqlStatement = fmt.Sprintf("INSERT INTO sensors (timestamp, average) VALUES (%+v, %+v)",list.Sensor1.Timestamp, avg)
	
	_, err = db.Exec(sqlStatement)
	if err != nil {
	  panic(err)
	}
}





func connectSensorDb() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  t.db = db

  fmt.Println("Successfully connected!")



}