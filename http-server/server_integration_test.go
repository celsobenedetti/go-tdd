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
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	playerName := "Celso"

	iw := httptest.NewRecorder() // ignore response writes
	server.ServeHTTP(iw, newPlayerReq(POST, playerName))
	server.ServeHTTP(iw, newPlayerReq(POST, playerName))
	server.ServeHTTP(iw, newPlayerReq(POST, playerName))

	t.Run("POST 3 wins and GET score", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newPlayerReq(GET, playerName))

		is.Equal(res.Body.String(), "3") // want 3 wins to be recorded for player
	})

	t.Run("POST 100 wins concurrently", func(t *testing.T) {
		server := NewPlayerServer(NewInMemoryPlayerStore())
		playerName := "Celso"

		wg := sync.WaitGroup{}
		wg.Add(100)

		for i := 0; i < 100; i++ {
			go func(name string) {
				server.ServeHTTP(iw, newPlayerReq(GET, playerName))
				wg.Done()
			}(playerName)
		}

		wg.Wait()
	})

	t.Run("GET /league", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newLeagueReq(GET))

		got, err := getLeagueFromResponse(t, res.Body)
		want := []Player{{"Celso", 3}}

		is.NoErr(err)       // unable to parse JSON
		is.Equal(want, got) // wanted different league JSON object
	})

}

func newPlayerReq(method string, playerName string) *http.Request {
	req, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", playerName), nil)
	return req
}

func newLeagueReq(method string) *http.Request {
	req, _ := http.NewRequest(method, "/league", nil)
	return req
}
