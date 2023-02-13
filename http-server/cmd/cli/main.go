package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/celso-patiri/go-tdd/http-server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	fmt.Println("Let's play Poker")
	fmt.Println("Type {Name} wins to record a victory")

	poker.NewCli(store, os.Stdin).PlayPoker()
}
