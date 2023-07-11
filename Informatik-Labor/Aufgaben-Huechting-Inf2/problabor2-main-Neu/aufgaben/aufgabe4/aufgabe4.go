// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe4

/* Aufgabenstellung:
 * In der Datei bintrree.go ist ein Datentyp für einen binären Suchbaum definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode HasAtLeastNElements().
 * Die Methode soll true liefern, falls der Baum mehr als N Elemente hat.
 */

// Liefert true, falls der Baum mindestens n Elemente hat.
func (tree *BinTree) HasAtLeastNElements(n int) bool {
	// TODO
	return tree.CountElements() >= n 
}


func (tree *BinTree) CountElements() int {
	if tree.Empty(){
		return 0 
	}
	return tree.Right.CountElements() + tree.Left.CountElements() +1
}