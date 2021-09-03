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
	SavePlayerScore(name string, score int)
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
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

const applicationJsonContentType = "application/json"

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", applicationJsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
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
