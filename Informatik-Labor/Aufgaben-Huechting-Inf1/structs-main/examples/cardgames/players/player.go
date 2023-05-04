package players

import (
	"fmt"

	"github.com/tel22a-inf/go.structs/examples/cardgames/cards"
)

/*
Player repräsentiert einen Spieler.

Ein Spieler ist zusammengesetzt aus einen Namen (string)
und einem Kartenstapel, seinen Handkarten.
Die Handkarten sind vom Typ Deck, weil ein Satz Handkarten
technisch gesehen nichts anderes ist als ein Kartenstapel.

Anmerkung: Wir geben den Handkarten keinen Namen, sondern verwenden nur den Typ Deck
innerhalb des Player-Structs. Dies ist ein sog. *struct embedding*.
Es hat den Vorteil, dass die Methoden von Deck direkt auch für Player zur Verfügung
stehen. Um einem Spieler Karten zu geben, müssen wir keine neue Methode schreiben,
sondern wir "recyclen" die Methode Add, die schon für Deck definiert ist.
*/
type Player struct {
	Name string
	cards.Deck
}

// New erzeugt einen neuen Spieler.
// Die Funktion erwartet den Namen eines Spielers und beliebig viele Karten.
// Sie konstruiert ein Player-Objekt mit diesem Namen und einem Satz Handkarten.
func New(name string, handcards ...cards.PlayingCard) Player {
	return Player{name, cards.NewDeck(handcards...)}
}

// String liefert eine menschenlesbare String-Repräsentation des Spielers.
func (player Player) String() string {
	if player.Length() == 0 {
		return fmt.Sprintf("%s: Keine Handkarten", player.Name)
	}
	return fmt.Sprintf("%s:\n%s", player.Name, player.Deck)
}
