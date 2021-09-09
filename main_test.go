package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"

    // "bytes"
    "strings"
)

func TestEventsHandler(t *testing.T) {
    r := strings.NewReader("{\"message\": \"Stay slothful!\"}\\n{\"message\": \"Stay slothful!\"}")

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