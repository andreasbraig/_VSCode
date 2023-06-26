// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe4

/* Aufgabenstellung:
 * In der Datei bintrree.go ist ein Datentyp für einen binären Suchbaum definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode HasOddElementNumber().
 * Die Methode soll true liefern, falls der Baum eine ungerade Anzahl an Elementen hat.
 */

func (tree *BinTree) HasOddElementNumber() bool {
	// TODO

	return (tree.CountElements() % 2) == 1
}

// Rekursive Funktion um die Elemente in einem Binären Suchbaum zu zählen
func (tree *BinTree) CountElements() int {
	if tree.Empty() {
		return 0
	}

	return tree.Left.CountElements() + tree.Right.CountElements() + 1

}
