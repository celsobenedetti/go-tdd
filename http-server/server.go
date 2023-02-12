package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

const (
	GET  = http.MethodGet
	POST = http.MethodPost
)

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	playerName := strings.TrimPrefix(r.URL.Path, "/players/")

	handler := func(cb func(http.ResponseWriter, string)) {
		cb(w, playerName)
	}

	switch r.Method {
	case GET:
		handler(s.showScore)
	case POST:
		handler(s.processWin)
	}

}

type PlayerServer struct {
	store PlayerStore
	mutex sync.Mutex
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
}

func NewPlayerServer() *PlayerServer {
    return &PlayerServer{NewInMemoryPlayerStore(), sync.Mutex{}}
}

func (s *PlayerServer) showScore(w http.ResponseWriter, playerName string) {
	score, ok := s.store.GetPlayerScore(playerName)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) processWin(w http.ResponseWriter, playerName string) {
	s.store.RecordWin(playerName)
	w.WriteHeader(http.StatusAccepted)
}
