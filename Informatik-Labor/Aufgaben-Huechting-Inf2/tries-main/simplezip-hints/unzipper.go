package simplezip

import (
	"github.com/tel22a-inf/tries/tries"
)

// Unzipper is a data structure that stores a compressed string and everything that is
// needed to track the progress while decompressing it.

type Unzipper struct {
	compressed string

	tw       *tries.TrieWalker
	position int

	result string
}

// NewUnzipper creates a new unzipper from a compressed string and a trie.
func NewUnzipper(compressed string, trie *tries.Trie) *Unzipper {
	return &Unzipper{
		tw:         tries.NewTrieWalker(trie),
		compressed: compressed,
	}
}

// Done returns true if the unzipper is done.
func (uz *Unzipper) Done() bool {
	// TODO
	return false

	// Der Unzipper ist fertig, wenn er am Ende der komprimierten Zeichenkette angekommen ist.
}

// Result returns the decompressed string.
func (uz *Unzipper) Result() string {
	return uz.result
}

// AdvanceSingleChar advances the unzipper by one character.
// It copies the current character to the result and advances the position.
//
// This function does not query the trie.
// It is meant to be used while unzipping a string when the trie walker has not moved.
func (uz *Unzipper) AdvanceSingleChar() {
	// TODO

	// Füge das aktuelle Zeichen der komprimierten Zeichenkette zum Ergebnis hinzu und
	// erhöhe die Position um eins.
	// Das "aktuelle Zeichen" ist das Zeichen an der Position in der komprimierten Zeichenkette.
	// Die Position ist der Index, an dem der Unzipper gerade steht.
}

// InsertCurrentData inserts the data of the current node into the result.
// If the current node has no data, it does nothing
// Returns true if data has been inserted.
//
// This function neither queries the trie nor advances the unzipper.
// It is meant to be used while unzipping a string after the trie walker has moved.
func (uz *Unzipper) InsertCurrentData() bool {
	// TODO
	return false

	// Hole die Daten des aktuellen Knotens.
	// Wenn der aktuelle Knoten Daten hat, füge die erste Datenkomponente zum Ergebnis hinzu.
	// - verwende z.B. fmt.Sprintf dafür.
	// - (Im Fall der Kompression mit simplezip ist das immer genau ein Element.)
	// Gib zurück, ob Daten eingefügt wurden.
}

// CopyChars expects a number of characters and copies them to from the compressed string to the result.
//
// This function neither queries the trie nor advances the unzipper.
// It is meant to be used while unzipping a string when the trie walker has not moved.
func (uz *Unzipper) CopyChars(n int) {
	// TODO

	// Berechne aufgrund von n und uz.position den Index des letzten zu kopierenden Zeichens.
	// Füge die Zeichen von uz.position bis zum letzten Zeichen zum Ergebnis hinzu.
}

// Run decompresses the string.
func (uz *Unzipper) Run() {
	// TODO

	// Rufe uz.tw.Walk() auf und speichere das Ergebnis in einer Variablen.
	// Diese Variable sagt uns, wie viele Zeichen der Trie Walker konsumiert hat.

	// Wenn der Trie Walker nichts konsumiert hat, rufe AdvanceSingleChar() auf.
	// Andernfalls rufe InsertCurrentData() auf, um die Daten des aktuellen Knotens einzufügen.
	// Wenn InsertCurrentData() nichts eingefügt hat, rufe CopyChars() auf, um die Zeichen zu kopieren.

	// Setze den Trie Walker zurück und erhöhe die Position um die Anzahl der konsumierten Zeichen.
	// Vorsicht: AdwanceSingleChar() erhöht die Position bereits um eins.

	// Wiederhole, bis der Unzipper fertig ist.
}
