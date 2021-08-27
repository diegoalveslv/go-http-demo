package main

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) SavePlayerScore(name string, score int) {
	s.scores[name] = score
}


type SpyPlayerStore struct {
	scores map[string]int
	nameStored  string
	scoreStored int
}

func (s *SpyPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *SpyPlayerStore) SavePlayerScore(name string, score int) {
	s.nameStored = name
	s.scoreStored = score
}