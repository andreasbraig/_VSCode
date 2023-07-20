// Thema Aufgabe: Listen
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe2

/* Aufgabenstellung:
 * In der Datei linkedlist.go  ist ein Datentyp für eine verkettete Liste definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode Filter().
 * Die Methode soll ein Prädikat (eine Bool-Funktion) Funktion erwarten.
 * Sie soll eine neue Liste liefern, die nur die Elemente aus der ursprünglichen Liste
 * enthält, für die das Prädikat true liefert.
 */

// Liefert eine Liste mit nur den Elementen, für die p true liefert.
func (list *LinkedList) Filter(p func(string) bool) *LinkedList {
	// TODO

	if list.Next == nil {
		return list
	}
	if p(list.Id) {

		list.Next = list.Next.Filter(p)
		return list
	}

	return list.Next.Filter(p)

}
