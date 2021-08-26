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
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
	}

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
}
