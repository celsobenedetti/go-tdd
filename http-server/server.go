package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// PlayerStore stores information about players
type PlayerStore interface {
	GetPlayerScore(playerName string) int
	RecordWin(playerName string)
	GetLeague() League
}

// Player stores a player name and the number of wins
type Player struct {
	Name string
	Wins int
}

// PlayerServer is an HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	mutex *sync.Mutex
	http.Handler
}

// NewPlayerServer creates a PlayerServer with routing configured
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
	w.Header().Add("content-type", jsonContentType)
	json.NewEncoder(w).Encode(s.store.GetLeague())
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
	score := s.store.GetPlayerScore(playerName)

	if score == 0 {
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

func (s *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Celso", 20},
		{"Joao", 10},
	}
}

const (
	GET             = http.MethodGet
	POST            = http.MethodPost
	jsonContentType = "application/json"
)
