package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Appolo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {

	t.Run("it returns accepted on POST", func(t *testing.T) {
		store := StubPlayerStore{
			map[string]int{},
		}
		server := &PlayerServer{&store}

		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it records wins when POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		storeSpy := SpyPlayerStore{
			scores: map[string]int{
				"Pepper": 20,
			},
		}
		server := &PlayerServer{&storeSpy}

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		assertWasStored(t, storeSpy, "Pepper", 21)
	})
}

func assertWasStored(t *testing.T, spy SpyPlayerStore, name string, score int) {
	t.Helper()

	if spy.nameStored != name {
		t.Errorf("incorrect name stored. Got %q want %q", spy.nameStored, name)
	}

	if spy.scoreStored != score {
		t.Errorf("incorrect score stored. Got %d want %d", spy.scoreStored, score)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get correct status, got %d want %d", got, want)
	}
}
