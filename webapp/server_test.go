package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListenAndServe(t *testing.T) {

	request, _ := http.NewRequest(http.MethodGet, "players/Pepper", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)

	got := strings.TrimSuffix(response.Body.String(), "\n")

	want := "20"

	if got != want {

		t.Errorf("got %q but want %q", got, want)

	}

}
