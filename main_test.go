package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"

    // "bytes"
    "time"
    "strings"
    "encoding/json"

    "log"

    faker "github.com/brianvoe/gofakeit/v6"
)

const (
    timestampFormat = "2006-01-02 15:04:05"
)

type TestEvent struct {
	ClientTime 	string	    `json:"client_time"`
	DeviceId 	string 		`json:"device_id"`
	DeviceOs 	string 		`json:"device_os"`
	Session		string		`json:"session"`
	Sequence  	int8		`json:"sequence"`
	Event	 	string		`json:"event"`
	ParamInt 	int8		`json:"param_int"`
	ParamStr 	string		`json:"param_str"`
}

func TestEventsHandler(t *testing.T) {
    clientTime := faker.DateRange(time.Now().Add(-24*time.Hour), time.Now()).Format(timestampFormat)

    event := TestEvent{
        ClientTime: clientTime,
        DeviceId: faker.UUID(),
        DeviceOs: "iOS 13.5.1",
        Session: faker.LetterN(16),
        Sequence: 1,
        Event: faker.RandomString([]string{"app_start", "app_stop"}),
        ParamInt: faker.Int8(),
        ParamStr: faker.HipsterSentence(2),
    }

    eventJson, _ := json.Marshal(event)
    r := strings.NewReader(string(eventJson))

    log.Println(string(eventJson))

    req := httptest.NewRequest(http.MethodPost, "/foo", r)
	w := httptest.NewRecorder()

	eventsHandler(w, req)

    res := w.Result()
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("expected error to be nil got %v", err)
    }
    if string(data) != "{\"status\": \"ok\"}" {
        t.Errorf("expected status ok got %v", string(data))
    }
}