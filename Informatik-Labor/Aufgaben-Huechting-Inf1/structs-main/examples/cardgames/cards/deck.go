package cards

import "math/rand"

/*
Deck ist der Datentyp für einen Kartenstapel.

Ein Kartenstapel ist nichts anderes als eine Liste von Karten.
Wir definieren ihn als Struct, das eine Liste von Karten enthält.
Alternativ könnte man das Deck direkt als Liste von Karten definieren,
dann wären aber manche Methoden schwerer lesbar.
*/
type Deck struct {
	cards []PlayingCard
}

// String liefert eine menschenlesbare String-Repräsentation eines Kartenstapels.
// Dafür wird die String-Methode von PlayingCard.
func (deck Deck) String() string {
	result := ""
	for _, card := range deck.cards {
		result += card.String() + "\n"
	}
	return result
}

// Length liefert die Anzahl der Karten des Stapels.
func (deck Deck) Length() int {
	return len(deck.cards)
}

// Get erwartet eine Nummer und liefer die entsprechende Karte aus dem Stapel.
func (deck Deck) Get(n int) PlayingCard {
	return deck.cards[n]
}

// Add fügt die gegebenen Spielkarten zum Deck hinzu (ans Ende).
func (deck *Deck) Add(cards ...PlayingCard) {
	deck.cards = append(deck.cards, cards...)
}

// Remove entfernt die Karte mit der angegebenen Position.
// Die Funktion ist nicht sicher, falls die Position nicht existiert.
//
// Anmerkung: Der Receiver dieser Methode ist ein Pointer.
// D.h. die Methode bekommt nur die Adresse eines Kartenstapels,
// nicht den Stapel selbst. Dies ist notwendig, da der Stapel
// von der Methode verändert wird.
func (deck *Deck) Remove(n int) {
	deck.cards = append(deck.cards[:n], deck.cards[n+1:]...)
}

// DrawCard zieht eine Karte vom Stapel.
// D.h. entfernt eine Karte und liefert diese.
// Diese Funktion ist nicht sicher, wenn der Stapel leer ist.
func (deck *Deck) DrawCard() PlayingCard {
	result := deck.Get(deck.Length() - 1)
	deck.cards = deck.cards[:deck.Length()-1]
	return result
}

// Shuffle mischt einen Kartenstapel.
func (deck Deck) Shuffle() {
	// Wir benutzen die Funktion `Shuffle()`` aus dem Package math/rand.
	// Diese kann beliebige Listen durcheinander bringen.
	// Dazu benötigt sie die Länge der Liste sowie eine Funktion, die zwei Elemente
	// erwartet und diese vertauscht. Wir müssen also eine solche Funktion definieren
	// und dann an `Shuffle()` übergeben.

	// Definiere eine Funktion, die zwei beliebige Karten aus dem Deck vertauscht.
	swapCards := func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}

	// Verwende rand.Shuffle() mit unserer Tausch-Funktion, um den Stapel zu mischen.
	rand.Shuffle(deck.Length(), swapCards)
}

// NewDeck liefert einen Kartenstapel, der die angegebenen Spielkarten enthält.
func NewDeck(cards ...PlayingCard) Deck {
	return Deck{cards}
}

// NewDeck32 liefert einen "frischen" sortierten (Skat-)Kartenstapel mit 32 Karten.
func NewDeck32() Deck {
	suits := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	ranks := []string{"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	deck := Deck{}
	for _, c := range suits {
		for _, v := range ranks {
			deck.Add(PlayingCard{c, v})
		}
	}
	return deck
}
