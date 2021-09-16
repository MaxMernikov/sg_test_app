package main

import (
	"log"
	"io/ioutil"
	"net"
	"net/http"
	"bytes"
)

func eventsHandler(w http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	requestIP, err := getIP(request)
	if err != nil {
		panic(err)
	}

    // log.Println(string(data))
	CreateEventsBatch(splitData(data), requestIP)
    w.Write([]byte(`{"status": "ok"}`))
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

func splitData(data []byte) [][]byte {
    separator := []byte("\\n")
	return bytes.Split(data, separator)
}

func main() {
	http.HandleFunc("/", eventsHandler)
	http.ListenAndServe(":8090", nil)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
