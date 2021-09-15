package main

import (
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"encoding/json"
	"time"
)

type Event struct {
	Ip			string  	`db:"ip"`
	ServerTime  time.Time  	`db:"server_time"`
	ClientTime 	time.Time	`db:"client_time"`
	DeviceId 	string 		`db:"device_id"`
	DeviceOs 	string 		`db:"device_os"`
	Session		string		`db:"session"`
	Sequence  	int8		`db:"sequence"`
	Event	 	string		`db:"event"`
	ParamInt 	int8		`db:"param_int"`
	ParamStr 	string		`db:"param_str"`
}

var (
	connect, _ = sqlx.Open("clickhouse", "tcp://localhost:9000?username=&compress=true&database=saygames_test")
)

func CreateEventsBatch (events [][]byte, requestIP string) {
	for _, eventBytes := range events {
		event := Event{}
		json.Unmarshal(eventBytes, &event)

		log.Println(string(eventBytes))
		log.Println(event.Ip)
	}

	log.Println(requestIP)
}

func save (events []Event) {
	var (
		tx, _   = connect.Begin()
		stmt, _ = tx.Prepare("INSERT INTO events (ip, server_time, client_time, device_id, device_os, session, sequence, event, param_int, param_str) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	)
	defer stmt.Close()


}