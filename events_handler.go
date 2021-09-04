package main

import (
	// "fmt"
    "net"
	"net/http"
	"io/ioutil"
    // "log"
    "models"
)

func eventsHandler(w http.ResponseWriter, request *http.Request) {
    events, err := ioutil.ReadAll(request.Body)
    if err != nil {
        panic(err)
    }

    requestIP, err := getIP(request)
    if err != nil {
        panic(err)
	}

    go models.CreateEventsBatch(events, requestIP)
}

func getIP(request *http.Request) (string, error) {
    ip, _, err := net.SplitHostPort(request.RemoteAddr)
    if err != nil {
        return "", err
    }
    netIP := net.ParseIP(ip)
    if netIP != nil {
        return ip, nil
    }
    return "", nil
}

func main() {
	http.HandleFunc("/", eventsHandler)
	http.ListenAndServe(":8090", nil)
}