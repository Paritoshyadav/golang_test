package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StubPlayerSore struct {
	score    map[string]int
	winCalls []string
}

func (s *StubPlayerSore) GetPlayerScore(name string) int {

	return s.score[name]

}
func (s *StubPlayerSore) RecordWins(name string) {

	s.winCalls = append(s.winCalls, name)

}

func TestListenAndServe(t *testing.T) {

	playerServer := &PlayerServer{&StubPlayerSore{map[string]int{
		"Pepper": 20,
		"Sam":    10,
	}, nil}}

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
	const playerName = "Sam"
	store := &StubPlayerSore{map[string]int{}, nil}
	playerServer := &PlayerServer{store}
	t.Run("accepting Post request", func(t *testing.T) {
		request := NewPostWinRequest(playerName)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Code

		want := http.StatusAccepted
		// assertPlayerScore(t, got, want)

		assertResponseStatus(t, got, want)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d but want %d", len(store.winCalls), 1)

		}
		if store.winCalls[0] != playerName {
			t.Errorf("got %s but want %s", store.winCalls[0], playerName)

		}

	})

}

func TestAndRecordWinAndScore(t *testing.T) {

	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))

	response := httptest.NewRecorder()

	server.ServeHTTP(response, NewRequest(player))

	assertResponseStatus(t, response.Code, http.StatusAccepted)
	assertPlayerScore(t, strings.TrimSuffix(response.Body.String(), "\n"), "3")

}

func TestLeague(t *testing.T) {

	store := &StubPlayerSore{}
	server := &PlayerServer{store}
	request, _ := http.NewRequest(http.MethodGet, "/league/", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)
	assertResponseStatus(t, response.Code, http.StatusAccepted)

}

func NewPostWinRequest(name string) *http.Request {

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req

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
