package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
    store PlayerStore
    in *bufio.Scanner
}

func (cli *CLI) PlayPoker()  {
    userInput := cli.ReadLine()
    cli.store.RecordWin(extractWinner(userInput))
}

func (cli *CLI) ReadLine() string {
    cli.in.Scan()
    return cli.in.Text()
}

func extractWinner(userInput string) string {
    return strings.TrimSuffix(userInput, " wins")
}

func NewCli(store PlayerStore, in io.Reader) *CLI {
    return &CLI{store, bufio.NewScanner(in)}
}

