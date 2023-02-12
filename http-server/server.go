package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

const (
	GET  = http.MethodGet
	POST = http.MethodPost
)

func NewPlayerServer(store PlayerStore) *PlayerServer {
	s := new(PlayerServer)

	s.store = store
	s.mutex = &sync.Mutex{}

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(s.playersHandle))

	s.Handler = router

	return s
}

func (s *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(s.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

func (s *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Celso", 20},
		{"Joao", 10},
	}
}

func (s *PlayerServer) playersHandle(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.String(), "/players/")

	switch r.Method {
	case GET:
		s.showScore(w, playerName)
	case POST:
		s.processWin(w, playerName)
	}
}

func (s *PlayerServer) showScore(w http.ResponseWriter, playerName string) {
	score, ok := s.store.GetPlayerScore(playerName)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) processWin(w http.ResponseWriter, playerName string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.store.RecordWin(playerName)
	w.WriteHeader(http.StatusAccepted)
}

type PlayerServer struct {
	store PlayerStore
	mutex *sync.Mutex
	http.Handler
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
    GetLeague() []Player
}

type Player struct {
	Name string
	Wins int
}
