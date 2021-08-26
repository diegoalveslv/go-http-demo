package main

import (
	"log"
	"net/http"
)

//main.go
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return 123
}

func (i *InMemoryPlayerStore) SavePlayerScore(name string, score int) {}

func main() {
    server := &PlayerServer{&InMemoryPlayerStore{}}
    log.Fatal(http.ListenAndServe(":5000", server))
}