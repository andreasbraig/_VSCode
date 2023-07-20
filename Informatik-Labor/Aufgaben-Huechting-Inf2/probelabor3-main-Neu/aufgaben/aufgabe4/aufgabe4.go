// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe4

/* Aufgabenstellung:
 * In der Datei bintrree.go ist ein Datentyp für einen binären Suchbaum definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode HasEvenElement().
 * Die Methode soll true liefern, falls der Baum wenigstens ein Element mit einem
 * geraden Wert enthält.
 */

// Liefert true, falls es ein Element mit geradem Wert im Baum gibt.
func (tree *BinTree) HasEvenElement() bool {
	// TODO
	if tree.Empty() {
		return false
	}
	if CheckEven(tree.Value) {
		return true
	}

	return tree.Left.HasEvenElement() || tree.Right.HasEvenElement()
}

func CheckEven(n int) bool {
	return n%2 == 0
}
