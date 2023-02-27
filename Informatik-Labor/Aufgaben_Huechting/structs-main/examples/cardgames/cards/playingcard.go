package cards

import "fmt"

/*
PlayingCard ist ein Datentyp für Spielkarten.

Eine PlayingCard Enthält Felder für Farbe ("Suit") und Wert ("Rank").
Beides sind Strings, weil Farben Begriffe wie "Kreuz", "Pik" etc. sind und
weil Werte zwar Zahlen wie "7", "8", etc., aber auch Namen wie "Bube" sein können.
*/
type PlayingCard struct {
	Suit string
	Rank string
}

// NewCard Liefert eine neue Spielkarte mit der angegebenen Farbe und Wert.
func NewCard(suit, rank string) PlayingCard {
	return PlayingCard{Suit: suit, Rank: rank}
}

// String liefert eine menschenlesbare String-Repräsentation einer Spielkarte.
// Durch die Existenz dieser Funktion wird fmt.Println() anstelle der generischen
// Repräsentation eines Structs mit geschweiften Klammern diese Funktion verwenden.
func (card PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", card.Rank, card.Suit)
}
