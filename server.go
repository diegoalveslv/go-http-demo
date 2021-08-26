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

	player := p.getPlayerName(r)

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	w.WriteHeader(http.StatusAccepted)

	playerScore := p.store.GetPlayerScore(player)

	playerScore++
	p.store.SavePlayerScore(player, playerScore)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	playerScore := p.store.GetPlayerScore(player)

	if playerScore == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	fmt.Fprint(w, playerScore)
}

func (p *PlayerServer) getPlayerName(r *http.Request) string {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	return player
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	SavePlayerScore(name string, score int)
}
