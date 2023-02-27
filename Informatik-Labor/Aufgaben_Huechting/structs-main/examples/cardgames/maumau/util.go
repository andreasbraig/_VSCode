package maumau

import (
	"fmt"

	"github.com/tel22a-inf/go.structs/examples/cardgames/players"
)

// askMove fragt einen Spieler nach seinem gewünschten Zug und liefert die
// Position der entsprechenden Karte.
// Entfernt die Karte noch nicht aus der Hand des Spielers.
func askMove(p *players.Player) int {
	fmt.Println(p)
	fmt.Println()

	fmt.Print("Bitte eine Karte wählen: ")
	var cardNumber int
	fmt.Scanln(&cardNumber)

	return cardNumber - 1
}
