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

var (
	eventsChannel = make(chan Event, ClickhouseInsertBatchSize)
)

func init() {
	go BatchEventsManager()
}

func main() {
	http.HandleFunc("/", eventsHandler)
	http.ListenAndServe(":8090", nil)
}

func BatchEventsManager() {
	batches := BatchEvents(eventsChannel, BatchMaxItems, BatchMaxTimeout)
	for batch := range batches {
		save(batch)
	}
}

func eventsHandler(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	if request.Method == "POST" {
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println(err)
			render400(&w)

			return
		}

		requestIP, err := GetIP(request)
		if err != nil {
			log.Println(err)
		}

		events, err := BuildEvents(data, requestIP, time.Now())

		if err != nil {
			log.Println(err, "request:", string(data))
			render400(&w)

			return
		}

		for _, event := range events {
			eventsChannel <- event
		}
	}

	w.Write([]byte(`{"status": "ok"}`))
}

func render400(w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Write([]byte(`{"status": "fail", "message": "Invalid request parameters"}`))
}