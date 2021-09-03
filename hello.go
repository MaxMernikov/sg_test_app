package main

import (
	// "fmt"
    "net"
	"net/http"
	"io/ioutil"
    "log"
)

func hello(w http.ResponseWriter, request *http.Request) {
    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))

    requestIP, err  := getIP(request)
    if err != nil {
        panic(err)
	}
    log.Println(requestIP)
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
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}