package main

import (
	"log"
)

type Event struct {
	clientTime 	string	`db:"client_time"` 	//
	deviceId 	string 	`db:"device_id"`	// UUID
	deviceOs 	string 	`db:"device_os"`
	session		string	`db:"session"`
	sequence  	int		`db:"sequence"`
	event	 	string	`db:"event"`
	paramInt 	int		`db:"param_int"`
	paramStr 	string	`db:"param_str"`
}

// "client_time":"2020-12-01 23:59:00",
// "device_id":"0287D9AA-4ADF-4B37-A60F-3E9E645C821E"
// "device_os":"iOS 13.5.1",
// "session":"ybuRi8mAUypxjbxQ",
// "sequence":1,
// "event":"app_start",
// "param_int":0,
// "param_str":"some text"

func CreateEventsBatch (events [][]byte, requestIP string) {
	for _, event := range events {
		log.Println(string(event))
	}

	log.Println(requestIP)
}