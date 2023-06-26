// Thema der Aufgabe: Listen
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe1

/* Aufgabenstellung:
 * In der Datei linkedlist.go  ist ein Datentyp für eine verkettete Liste definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode SubList().
 * Die Methode soll zwei Positionen begin und end und die Teilliste liefern,
 * die zwischen diesen beiden Positionen liegt.
 * Genauer soll das Element an Position begin dazugehören und das Element an Postion
 * end soll das erste Element sein, das nicht mehr dazugehört.
 * Liegt eine der Positionen nicht in der Liste, soll die Teilliste in dieser Richtung
 * so lang wie möglich sein. Ist begin>=end, soll die leere Liste geliefert werden.
 */

// Liefert die Teilliste im halboffenen Intervall [begin, end[
func (list *LinkedList) SubList(begin, end int) *LinkedList {
	// TODO
	if end < begin {
		return nil 
	}

	if begin == 0 {

		el := list
		for i := 0 ; i <= end ; i++ {
			if i == end {
				el.Next = nil 
			}
			el = el.Next
		}


		return list 

	}

	return list.Next.SubList(begin-1, end-1)
}
