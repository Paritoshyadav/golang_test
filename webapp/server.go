package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		p.getScore(rw, req)
	case http.MethodPost:
		p.processWin(rw, req)

	}

}

func (p *PlayerServer) getScore(rw http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	if score == 0 {

		rw.WriteHeader(http.StatusNotFound)

	}

	fmt.Fprintln(rw, score)

}

func (p *PlayerServer) processWin(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusAccepted)

}
