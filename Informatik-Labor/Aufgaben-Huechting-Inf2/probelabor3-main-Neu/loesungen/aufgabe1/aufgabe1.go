// Thema der Aufgabe: Listen
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe1

/* Aufgabenstellung:
 * In der Datei linkedlist.go ist ein Datentyp für eine verkettete Liste definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode FirstDuplicate().
 * Die Methode soll die Position des ersten Elements liefern, das zwei Mal hintereinander
 * vorkommt. Gibt es kein solches Element, soll die Funktion die Länge der Liste liefern.
 */

// Liefert die erste Position, an der zwei gleiche Elemente hintereinander stehen.
func (list *LinkedList) FirstDuplicate() int {
	if list.Empty() {
		return 0
	}
	if list.Next.Empty() {
		return 1
	}
	if list.Id == list.Next.Id {
		return 0
	}
	return 1 + list.Next.FirstDuplicate()
}
