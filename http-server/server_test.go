package poker

import (
	"fmt"
	"net/http"
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
		nil,
	}
	server := NewPlayerServer(&playerStore)

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
		req, res := newRequestAndResponse(http.MethodGet, "/players/NIL")

		server.ServeHTTP(res, req)

		is.Equal(res.Code, http.StatusNotFound) // want status to be 404
	})

}

func TestScoreWins(t *testing.T) {
	is := is.New(t)

	playerStore := StubPlayerStore{
		map[string]int{},
		[]string{},
		nil,
	}
	server := NewPlayerServer(&playerStore)

	t.Run("it records win on POST", func(t *testing.T) {
		player := "Celso"
		req, res := newRequestAndResponse(http.MethodPost, fmt.Sprintf("/players/%s", player))

		server.ServeHTTP(res, req)

		is.Equal(len(playerStore.WinCalls), 1)      // want 1 win to be recorded
		is.Equal(res.Code, http.StatusAccepted)     // want status to be 202
		is.True(playerStore.hasRecordedWin(player)) // want player win to be recorded
	})
}

func TestLeague(t *testing.T) {
	is := is.New(t)

	t.Run("it returns league table as JSON", func(t *testing.T) {
		wantedLeague := League{
			{"Celso", 20},
			{"Joao", 10},
		}

		playerStore := StubPlayerStore{nil, nil, wantedLeague}

		server := NewPlayerServer(&playerStore)
		req, res := newRequestAndResponse(http.MethodGet, "/league")

		server.ServeHTTP(res, req)

		got, err := ParseLeagueJSON(res.Body)

		is.NoErr(err)                                                      // unable to parse JSON
		is.Equal(res.Result().Header.Get("content-type"), jsonContentType) // should heave application/json header
		is.Equal(wantedLeague, got)                                        // wanted different league JSON object
		is.Equal(res.Code, http.StatusOK)                                  // wanted status code to be 200
	})
}

