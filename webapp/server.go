package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWins(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/league/", http.HandlerFunc(p.leagueHandler))

	router.ServeHTTP(rw, req)

}
func (p *PlayerServer) leagueHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusAccepted)
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
