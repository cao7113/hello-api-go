package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	ping(w, req)

	res := w.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	var ping Ping

	err = json.Unmarshal(data, &ping)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if ping.Status != "pong" {
		t.Errorf("expected status to be pong got %v", err)
	}
}
