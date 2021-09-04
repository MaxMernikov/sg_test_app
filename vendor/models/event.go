package models

import (
	"log"
    "bytes"
)

func CreateEventsBatch (events []byte, requestIP string) {
    separator := []byte("\\n")
    events_ := bytes.Split(events, separator)

    for _, event := range events_ {
        log.Println(string(event))
    }

    log.Println(requestIP)
}