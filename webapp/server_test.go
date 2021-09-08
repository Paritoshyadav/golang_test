package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type StubPlayerSore struct {
	score    map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerSore) GetPlayerScore(name string) int {

	return s.score[name]

}
func (s *StubPlayerSore) RecordWins(name string) {

	s.winCalls = append(s.winCalls, name)

}

func (s *StubPlayerSore) GetLeague() []Player {

	return s.league

}

func TestListenAndServe(t *testing.T) {
	store := &StubPlayerSore{map[string]int{
		"Pepper": 20,
		"Sam":    10,
	}, nil, nil}
	playerServer := NewPlayerServer(store)

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
	store := &StubPlayerSore{map[string]int{}, nil, nil}
	playerServer := NewPlayerServer(store)
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
	server := NewPlayerServer(store)
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
	t.Run("checking if json response is serving", func(t *testing.T) {

		store := &StubPlayerSore{}
		server := NewPlayerServer(store)
		request, _ := http.NewRequest(http.MethodGet, "/league/", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}
		if response.Result().Header.Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
		}
		assertResponseStatus(t, response.Code, http.StatusOK)

	})

	t.Run("it return a league table", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := &StubPlayerSore{nil, nil, wantedLeague}

		server := NewPlayerServer(store)
		request, _ := http.NewRequest(http.MethodGet, "/league/", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}
		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v but wanted %v", got, wantedLeague)
		}

	})

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
