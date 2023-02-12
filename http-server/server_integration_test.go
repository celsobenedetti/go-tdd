package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/matryer/is"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	is := is.New(t)

	t.Run("POST 3 wins and GET score", func(t *testing.T) {
		server := &PlayerServer{NewInMemoryPlayerStore(), sync.Mutex{}}
		playerName := "Celso"

		i := httptest.NewRecorder() // ignore response writes
		server.request(http.MethodPost, playerName, i)
		server.request(http.MethodPost, playerName, i)
		server.request(http.MethodPost, playerName, i)

		res := httptest.NewRecorder()
		server.request(http.MethodGet, playerName, res)

		is.Equal(res.Body.String(), "3") // want 3 wins to be recorded for player
	})

	t.Run("POST and GET 100 wins concurrently", func(t *testing.T) {
		server := &PlayerServer{NewInMemoryPlayerStore(), sync.Mutex{}}
		playerName := "Celso"

	       wg := sync.WaitGroup{}
	       wg.Add(100)

		for i := 0; i < 100; i++ {
			go func(name string) {
				res := httptest.NewRecorder()
	               server.request(http.MethodGet, playerName, res)
	               want := res.Body.String()

	               server.request(http.MethodPost, playerName, httptest.NewRecorder())

	               res = httptest.NewRecorder()
	               server.request(http.MethodGet, playerName, res)
	               got := res.Body.String()

	               is.Equal(got, want)

	               wg.Done()
			}(playerName)
		}
	       wg.Wait()
	})

}

func (s *PlayerServer) request(method, playerName string, w http.ResponseWriter) {
	s.ServeHTTP(w, newPlayerReq(method, playerName))
}

func newPlayerReq(method string, playerName string) *http.Request {
	req, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", playerName), nil)
	return req
}
