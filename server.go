package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)

	player := strings.TrimPrefix(r.URL.Path, "/players/")
	playerScore := p.store.GetPlayerScore(player)

	playerScore += 1
	p.store.SavePlayerScore(player, playerScore)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	playerScore := p.store.GetPlayerScore(player)

	if playerScore == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	fmt.Fprint(w, playerScore)
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	SavePlayerScore(name string, score int)
}
