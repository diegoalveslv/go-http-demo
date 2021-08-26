package main

import (
	"log"
	"net/http"
)

//main.go
func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
