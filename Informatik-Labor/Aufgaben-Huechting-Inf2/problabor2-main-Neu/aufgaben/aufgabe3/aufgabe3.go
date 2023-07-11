// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe3

/* Aufgabenstellung:
 * In der Datei bintrree.go ist ein Datentyp für einen binären Suchbaum definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode DepthOf().
 * Die Methode soll einen Wert erwarten und die Tiefe dieses Knotens liefern.
 * Die Wurzel hat dabei die Tiefe 0, deren Kinder die Tiefe 1, die Enkel Tiefe 2 usw.
 * Gibt es keinen Knoten mit diesem Wert, soll -1 geliefert werden.
 *
 */

// Liefert die Tiefe des Knotens mit dem Wert value.
func (tree *BinTree) DepthOf(value int) int {
	// TODO
	if !tree.LookUp(value) {
		return -1
	}
	if value == tree.Value {
		return 0
	}
	if value > tree.Value {
		return tree.Right.DepthOf(value) + 1
	} else {
		return tree.Left.DepthOf(value) + 1
	}
}

func (tree *BinTree) LookUp(value int) bool {
	if tree.Empty() {
		return false
	}
	if value == tree.Value {
		return true
	}
	if value > tree.Value {
		return tree.Right.LookUp(value)
	} else {
		return tree.Right.LookUp(value)
	}

}
