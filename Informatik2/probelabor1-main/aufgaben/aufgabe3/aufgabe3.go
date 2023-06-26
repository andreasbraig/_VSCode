// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe3

/* Aufgabenstellung:
 * In der Datei bintrree.go ist ein Datentyp für einen binären Suchbaum definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode DecreaseSmallest().
 * Die Methode soll herausfinden, welches der kleinste Wert im Baum ist.
 * Wenn x der kleinste Wert ist, soll die Methode ihn durch x-1 ersetzen.
 * Wenn der Baum leer ist, soll nichts geschehen.
 */

// Verringert den kleinsten Wert im Baum um 1.
func (tree *BinTree) DecreaseSmallest() {
	// TODO

	el := tree.FindSmallest()

	if el != nil {
		el.Value -= 1
	}

}

func (tree *BinTree) FindSmallest() *BinTree {

	if tree.Empty() {
		return nil
	}

	if tree.Left.Empty() {
		return tree
	}

	return tree.Left.FindSmallest()

}
