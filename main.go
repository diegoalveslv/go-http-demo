package main

import (
	"log"
	"net/http"
)

//main.go
type InMemoryPlayerStore struct {
	scores map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (i *InMemoryPlayerStore) SavePlayerScore(name string, score int) {
	i.scores[name] = score
}

func main() {
	scores := make(map[string]int)
	server := &PlayerServer{&InMemoryPlayerStore{scores}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
