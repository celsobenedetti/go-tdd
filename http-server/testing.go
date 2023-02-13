package poker

import (
	"net/http"
	"net/http/httptest"
)

// Internals

func newRequestAndResponse(method, path string) (req *http.Request, res *httptest.ResponseRecorder) {
	req, _ = http.NewRequest(method, path, nil)
	res = httptest.NewRecorder()
	return req, res
}

type StubPlayerStore struct {
	scores   map[string]int
	WinCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func (s *StubPlayerStore) hasRecordedWin(name string) bool {
	for _, winner := range s.WinCalls {
		if winner == name {
			return true
		}
	}
	return false
}
