package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/matryer/is"
)

func TestGETPlayers(t *testing.T) {
	is := is.New(t)

	playerStore := StubPlayerStore{
		map[string]int{
			"Celso": 20,
			"Joao":  10,
		},
		[]string{},
	}
	server := &PlayerServer{&playerStore, sync.Mutex{}}

	t.Run("get Celso's score", func(t *testing.T) {
		req, res := newRequestAndResponse(http.MethodGet, "/players/Celso")

		server.ServeHTTP(res, req)

		is.Equal(res.Body.String(), "20") // response body is wrong
		is.Equal(res.Code, http.StatusOK) // want status to be 200
	})

	t.Run("get Joao's score", func(t *testing.T) {
		req, res := newRequestAndResponse(http.MethodGet, "/players/Joao")

		server.ServeHTTP(res, req)

		is.Equal(res.Body.String(), "10") // response body is wrong
		is.Equal(res.Code, http.StatusOK) // want status to be 200
	})

	t.Run("return 404 on missing player", func(t *testing.T) {
		req, res := newRequestAndResponse(http.MethodGet, "players/NIL")

		server.ServeHTTP(res, req)

		is.Equal(res.Code, http.StatusNotFound) // want status to be 404
	})

}

func TestScoreWins(t *testing.T) {
	is := is.New(t)

	playerStore := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{&playerStore, sync.Mutex{}}

	t.Run("it records win on POST", func(t *testing.T) {
		player := "Celso"
		req, res := newRequestAndResponse(http.MethodPost, fmt.Sprintf("/players/%s", player))

		server.ServeHTTP(res, req)

		is.Equal(len(playerStore.winCalls), 1)  // want 1 win to be recorded
		is.Equal(res.Code, http.StatusAccepted) // want status to be 202
		is.True(playerStore.hasRecordedWin(player)) // want player win to be recorded
	})
}

func newRequestAndResponse(method, path string) (req *http.Request, res *httptest.ResponseRecorder) {
	req, _ = http.NewRequest(method, path, nil)
	res = httptest.NewRecorder()
	return req, res
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := s.scores[name]
	return score, ok
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) hasRecordedWin(name string) bool {
	for _, winner := range s.winCalls {
		if winner == name {
			return true
		}
	}
	return false
}
