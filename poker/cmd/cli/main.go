package main

import (
	"fmt"
	"github.com/edreg/awesome/poker"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	cli.PlayPoker()
}
