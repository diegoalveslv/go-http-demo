package main

//in_memory_player_store.go
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
    return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
    scores map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (i *InMemoryPlayerStore) SavePlayerScore(name string, score int) {
	i.scores[name] = score
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	return nil
}
