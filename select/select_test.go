package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {

	site1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)

		w.WriteHeader(http.StatusOK)

	}))

	site2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

	}))

	slowUrl := site1.URL
	fastUrl := site2.URL

	got := race(slowUrl, fastUrl)

	want := fastUrl

	if got != want {
		t.Errorf("got %v , want %v", got, want)
	}

}
