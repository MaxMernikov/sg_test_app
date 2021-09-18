package main

import (
	"bytes"
	"log"
	"net"
	"net/http"
)

func SplitRequestData(data []byte) [][]byte {
	separator := []byte("\n")
	return bytes.Split(data, separator)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetIP(request *http.Request) (string, error) {
	unknownIp := "0.0.0.0"

	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		return unknownIp, err
	}

	netIP := net.ParseIP(ip)
	if ip == "::1" {
		return "127.0.0.1", nil
	}
	if netIP != nil {
		return ip, nil
	}
	return unknownIp, nil
}