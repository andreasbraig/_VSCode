package simplezip

import "github.com/tel22a-inf/tries/tries"

// Decompress expects a compressed string and returns the decompressed string.
// It also expects a Trie that is used as a dictionary.
func Decompress(compressedString string, trie *tries.Trie) string {
	// TODO
	uz := NewUnzipper(compressedString, trie) //Erstellen eines neuen Unzippers
	uz.Run()                                  //Ausführen des Runs
	return uz.result                          //Rückgabe des Results
}
