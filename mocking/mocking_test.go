package mocking

import (
	"reflect"
	"testing"
)

const (
	sleep = "sleep"
	write = "write"
)

func TestCountdown(t *testing.T) {
	sleeper := spyCountdownCalls{}

	Countdown(&sleeper, &sleeper)

	countdownCallsWant := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if reflect.DeepEqual(sleeper.buffer, countdownCallsWant) {
		t.Errorf("Wrong callstack, got %v, wanted %v", sleeper.calls, countdownCallsWant)
	}

	gotMessage := string(sleeper.buffer)
	wantMessage := `3
2
1
Go!
`

	if gotMessage != wantMessage {
		t.Errorf("Wrong countdown message: got %q, wanted %q", gotMessage, wantMessage)
	}
}

type spyCountdownCalls struct {
	calls  []string
	buffer []byte
}

func (s *spyCountdownCalls) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *spyCountdownCalls) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	s.buffer = append(s.buffer, p...)
	return
}
