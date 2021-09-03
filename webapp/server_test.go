package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StubPlayerSore struct {
	score map[string]int
}

func (s *StubPlayerSore) GetPlayerScore(name string) int {

	return s.score[name]

}

func TestListenAndServe(t *testing.T) {

	playerServer := &PlayerServer{&StubPlayerSore{map[string]int{
		"Pepper": 20,
		"Sam":    10,
	}}}

	t.Run("return Pepper score", func(t *testing.T) {
		request := NewRequest("Pepper")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := strings.TrimSuffix(response.Body.String(), "\n")

		want := "20"

		assertPlayerScore(t, got, want)
	})

	t.Run("return Sam score", func(t *testing.T) {
		request := NewRequest("Sam")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := strings.TrimSuffix(response.Body.String(), "\n")

		want := "10"
		assertPlayerScore(t, got, want)

	})
	t.Run("return 404 missing player", func(t *testing.T) {
		request := NewRequest("Queen")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Code

		want := http.StatusNotFound
		// assertPlayerScore(t, got, want)

		assertResponseStatus(t, got, want)

	})

}

func TestStoreWins(t *testing.T) {
	playerServer := &PlayerServer{&StubPlayerSore{map[string]int{}}}
	t.Run("accepting Post request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/m", nil)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Code

		want := http.StatusAccepted
		// assertPlayerScore(t, got, want)

		assertResponseStatus(t, got, want)

	})

}

func NewRequest(name string) *http.Request {

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req

}

func assertResponseStatus(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {

		t.Errorf("got %d but want %d", got, want)

	}

}

func assertPlayerScore(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {

		t.Errorf("got %q but want %q", got, want)

	}

}
