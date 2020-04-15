package poker_test

import (
	poker "github.com/edreg/awesome/app"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := poker.NewStubPlayerStore(nil, nil, nil)

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := poker.NewStubPlayerStore(nil, nil, nil)

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

}