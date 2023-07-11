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
	// TODO

	el := list
	count := 0

	for el.Next != nil {
		if el.Id == el.Next.Id {
			return count
		}
		el = el.Next
		count++
	}

	return count
}
