package poker_test

import (
	"strings"
	"testing"

	poker "github.com/celso-patiri/go-tdd/http-server"
	"github.com/matryer/is"
)

func TestCLI(t *testing.T) {
	is := is.New(t)

	t.Run("record 1 win for Celso", func(t *testing.T) {
		in := strings.NewReader("Celso wins\n")

		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCli(playerStore, in)

		cli.PlayPoker()

		is.Equal(len(playerStore.WinCalls), 1)          // expected 1 win call
		is.Equal(playerStore.WinCalls[0], "Celso") // didn't record correct winner
	})

	t.Run("record 1 win for Joao", func(t *testing.T) {
		in := strings.NewReader("Joao wins\n")

		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCli(playerStore, in)
		cli.PlayPoker()

		is.Equal(len(playerStore.WinCalls), 1)         // expected 1 win call
		is.Equal(playerStore.WinCalls[0], "Joao") // didn't record correct winner
	})

}
