package maumau

import (
	"fmt"

	"github.com/tel22a-inf/go.structs/examples/cardgames/cards"
	"github.com/tel22a-inf/go.structs/examples/cardgames/game"
	"github.com/tel22a-inf/go.structs/examples/cardgames/players"
)

/*
Game ist ein Datentyp, der alle Daten für ein Mau-Mau-Spiel verwaltet.
*/
type Game struct {
	Players     []players.Player
	Deck        cards.Deck
	DiscardPile cards.Deck
}

// New erwartet eine Liste von Spielernamen als Strings.
// Die Funktion konstruiert ein Game-Objekt, das einen frischen Kartenstapel,
// einen leeren Ablagestapel sowie eine Liste von Player-Objekten enthält.
func New(playerNames ...string) Game {
	result := Game{
		Deck:        cards.NewDeck32(),
		DiscardPile: cards.NewDeck(),
		Players:     []players.Player{},
	}
	for _, name := range playerNames {
		result.Players = append(result.Players, players.New(name))
	}
	return result
}

// Setup initialisiert das Spiel gemäß den Mau-Mau-Regeln.
// D.h. die Funktion gibt jedem Spieler 6 Karten und
// zieht anschließend eine Karte auf den Ablagestapel.
func (g *Game) Setup() {
	for i := range g.Players {
		game.Deal(&g.Deck, &(g.Players[i]), 6)
	}
	g.DiscardPile.Add(g.Deck.DrawCard())
}

// Run startet das Spiel und läuft, bis es beendet ist.
func (g *Game) Run() {
	currentPlayerNumber := 0

	for g.Running() {
		currentPlayer := &g.Players[currentPlayerNumber]
		cardNumber := askMove(currentPlayer)
		g.DiscardPile.Add(currentPlayer.Get(cardNumber))
		currentPlayer.Remove(cardNumber)
		currentPlayerNumber = (currentPlayerNumber + 1) % len(g.Players)
	}
}

// Winner liefert den Spieler, der gewonnen hat.
// Liefert nil, falls das Spiel noch läuft.
func (g Game) Winner() *players.Player {
	for _, p := range g.Players {
		if p.Length() == 0 {
			return &p
		}
	}
	return nil
}

// Running liefert true, falls das Spiel (noch) läuft.
func (g Game) Running() bool {
	return g.Winner() == nil
}

// PrintState gibt das Spielergebnis auf der Konsole aus.
func (g Game) PrintState() {
	winner := g.Winner()
	if winner == nil {
		fmt.Println("Das Spiel läuft noch.")
	} else {
		fmt.Printf("%s hat gewonnen.\n", winner.Name)
	}
}
