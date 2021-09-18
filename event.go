package main

import (
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"encoding/json"
	"time"

	// "log"
	// "reflect"
)

type Event struct {
	Ip			string  	`db:"ip"`
	ServerTime  time.Time  	`db:"server_time"`
	ClientTime 	Timestamp	`db:"client_time" json:"client_time"`
	DeviceId 	string 		`db:"device_id" json:"device_id"`
	DeviceOs 	string 		`db:"device_os" json:"device_os"`
	Session		string		`db:"session"`
	Sequence  	int8		`db:"sequence"`
	Event	 	string		`db:"event"`
	ParamInt 	int8		`db:"param_int" json:"param_int"`
	ParamStr 	string		`db:"param_str" json:"param_str"`
}

func BuildEvents(data []byte, requestIP string, serverTime time.Time) ([]Event, error) {
	var events []Event
	eventsBytes := SplitRequestData(data)

	for _, eventBytes := range eventsBytes {
		event := Event{
			Ip: requestIP,
			ServerTime: serverTime,
		}

		err := json.Unmarshal(eventBytes, &event)
		CheckErr(err)

		events = append(events, event)
	}

	return events, nil
}

func BatchEvents(values <-chan Event, maxItems int, maxTimeout time.Duration) chan []Event {
	batches := make(chan []Event)

	go func() {
		defer close(batches)

		for keepGoing := true; keepGoing; {
			var batch []Event
			expire := time.After(maxTimeout)
			for {
				select {
				case value, ok := <-values:
					if !ok {
						keepGoing = false
						goto done
					}

					batch = append(batch, value)
					if len(batch) == maxItems {
						goto done
					}

				case <-expire:
					goto done
				}
			}

		done:
			if len(batch) > 0 {
				batches <- batch
			}
		}
	}()

	return batches
}

func save (events []Event) {
	connect, _ := sqlx.Open("clickhouse", "tcp://localhost:9000?username=&compress=true&database=saygames_test")

	tx, _   := connect.Beginx()
	stmt, _ := tx.PrepareNamed("INSERT INTO events (ip, server_time, client_time, device_id, device_os, session, sequence, event, param_int, param_str) VALUES (:ip, :server_time, :client_time, :device_id, :device_os, :session, :sequence, :event, :param_int, :param_str)")
	defer stmt.Close()

	for _, event := range events {
		_, err := stmt.Exec(event)
		CheckErr(err)
	}

	err := tx.Commit()
	CheckErr(err)
}