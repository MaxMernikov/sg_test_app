package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"log"
	_ "net/http/pprof"
)

const (
	ClickhouseInsertBatchSize = 2000000
	BatchMaxItems = 10000
	BatchMaxTimeout = 1 * time.Second
)

var eventsChannel = make(chan Event, ClickhouseInsertBatchSize)

func BatchEventsManager() {
	batches := BatchEvents(eventsChannel, BatchMaxItems, BatchMaxTimeout)
	for batch := range batches {
		log.Println("save", len(batch))
		save(batch)
	}
}

func eventsHandler(w http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	CheckErr(err)

	requestIP, err := GetIP(request)
	CheckErr(err)

	events, err := BuildEvents(data, requestIP, time.Now())
	CheckErr(err)

	for _, event := range events {
		eventsChannel <- event
	}

	events = nil

	w.Write([]byte(`{"status": "ok"}`))
	request.Header.Set("Connection", "close")
}

func main() {
	go BatchEventsManager()
	http.HandleFunc("/t", eventsHandler)
	http.ListenAndServe(":8090", nil)
}
