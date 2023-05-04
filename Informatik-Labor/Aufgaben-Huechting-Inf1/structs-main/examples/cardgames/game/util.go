package game

import (
	"github.com/tel22a-inf/go.structs/examples/cardgames/cards"
	"github.com/tel22a-inf/go.structs/examples/cardgames/players"
)

// Deal erwartet ein Deck, einen Spieler und eine Kartenzahl.
// Die Funktion entfernt diese Anzahl Karten vom Stapel und gibt sie dem Spieler.
func Deal(deck *cards.Deck, player *players.Player, count int) {
	for i := 0; i < count; i++ {
		player.Add(deck.DrawCard())
	}
}
