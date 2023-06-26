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
	// der Unzipper ist fertig, wenn er am Ende der komprimierten Zeichenkette ist
	return uz.position == len(uz.compressed)
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

	//Das Zeichen aus compressed an stelle der Current Position wird an das Result angehängt
	uz.result += string(uz.compressed[uz.position])

	//Die Position wird um eins erhöht
	uz.position++

}

// InsertCurrentData inserts the data of the current node into the result.
// If the current node has no data, it does nothing
// Returns true if data has been inserted.
//
// This function neither queries the trie nor advances the unzipper.
// It is meant to be used while unzipping a string after the trie walker has moved.
func (uz *Unzipper) InsertCurrentData() bool {
	// TODO

	//extrahiere die aktuellen Daten aus dem aktuellen knoten des Walkers
	data := uz.tw.Data()

	//Data ist eine Liste von Any
	//Wenn diese Liste Null ist, ist der Knoten leer wenn nicht kann etwas angefügt werden.
	if len(data) != 0 {

		//Da Data ja eine Liste ist, muss jedes Element nacheinander angefügt werden. IDR Läuft die schleife nur ein mal
		for i := range data {

			uz.result += data[i].(string) // Data ist ja Any und hier wird der Typ festgelegt, damit die Strings zusammengeführt werden.
		}
		return true

	}

	return false
}

// CopyChars expects a number of characters and copies them from the compressed string to the result.
//
// This function neither queries the trie nor advances the unzipper.
// It is meant to be used while unzipping a string when the trie walker has not moved.
func (uz *Unzipper) CopyChars(n int) {
	// TODO

	//Der Char wird N Mal an Result gehängt.
	for i := 0; i < n; i++ {
		uz.AdvanceSingleChar()
	}

}

// Run decompresses the string.
func (uz *Unzipper) Run() {
	// TODO

	//Solange der Unzipper nicht fertig ist 
	for !uz.Done() {
		//Berechne wie viele Buchstaben der Walker in den Baum gelaufen ist 
		consumed := uz.tw.Walk(uz.compressed[uz.position:])

		//Wenn er nicht gelaufen ist, kopiere den einen buchstaben weiter an Result
		if consumed == 0 {
			uz.AdvanceSingleChar()
			continue//und geh aus der if bedingung raus
		} else if !uz.InsertCurrentData() { //führt Insert Current Data aus und t 
			uz.CopyChars(consumed) //wenn es fehlschlägt werden alle gelaufenen buchstaben angehäng
		}
		uz.tw.Reset() //Der Walker wird zurückgesetzt, dass current=trie
		uz.position += consumed // Die Position wird um die gelaufenen Buchstaben erhöht 

	}

}
