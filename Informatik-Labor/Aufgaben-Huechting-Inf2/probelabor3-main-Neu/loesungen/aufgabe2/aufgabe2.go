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
	if list.Empty() {
		return MakeLinkedList()
	}
	filtered := list.Next.Filter(p)
	if !p(list.Id) {
		return filtered
	}
	result := MakeLinkedList(list.Id)
	result.Next = filtered
	return result
}
