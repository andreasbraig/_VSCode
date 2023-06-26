// Thema der Aufgabe: Listen
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe1

/* Aufgabenstellung:
 * In der Datei linkedlist.go  ist ein Datentyp für eine verkettete Liste definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode Erase().
 * Die Methode soll eine Position erwarten und aus der Liste das Element an dieser
 * Stelle entfernen. Kommt das Element nicht vor, soll die Funktion nichts machen.
 */

// Entfernt das Element mit der gegebenen Position aus der Liste.
func (list *LinkedList) Erase(pos int) {
	// TODO
	if list == nil {
		return
	}
	if pos == 0 {
		list = list.Next
		return
	}
	if pos == 1 {
		list.Next = list.Next.Next
	}
	if pos > 1 {
		list.Next.Erase(pos - 1)

	}
}

func SwapValues(e1, e2 *LinkedList) {

	e2.Value, e1.Value = e1.Value, e2.Value
}
