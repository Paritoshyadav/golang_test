package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWins(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()

	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/league/", http.HandlerFunc(p.leagueHandler))
	p.Handler = router

	return p

}

func (p *PlayerServer) leagueHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	err := json.NewEncoder(rw).Encode(p.getLeagueTable())

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func (p *PlayerServer) getLeagueTable() []Player {

	return p.store.GetLeague()

}

func (p *PlayerServer) playerHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getScore(rw, r)
	case http.MethodPost:
		p.processWin(rw, r)

	}

}

func (p *PlayerServer) getScore(rw http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	if score == 0 {

		rw.WriteHeader(http.StatusNotFound)

	}
	rw.WriteHeader(http.StatusAccepted)

	fmt.Fprintln(rw, score)

}

func (p *PlayerServer) processWin(rw http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	rw.WriteHeader(http.StatusAccepted)
	p.store.RecordWins(player)

}
