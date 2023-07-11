// Thema Aufgabe: Listen
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe2

/* Aufgabenstellung:
 * In der Datei linkedlist.go  ist ein Datentyp für eine verkettete Liste definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode Reverse().
 * Die Methode soll die Liste umdrehen und den neuen Head liefern.
 * D.h. die Methode soll die Elemente erhalten, die Pointer aber so umbiegen, dass
 * die Liste umgedreht ist.
 *
 * Anmerkung: Bei dieser Aufgabe würde in der Bewertung auch geprüft werden,
 * ob wirklich nur Pointer verdreht wurden, d.h. ob die Adressen der Elemente noch
 * die gleichen sind. Diese Forderung ist aber relativ leicht einzuhalten,
 * indem man einfach keine neuen Elemente erzeugt ;-).
 */

// Dreht die Liste um und liefert den neuen Kopf.
func (list *LinkedList) Reverse() *LinkedList {
	if list.Empty(){ //Grundsätzlich Check: Liste leer?
		return nil
	}
	if list.Next.Next == nil { //Rekursionsanker Abbruch wenn nur noch ein Element
		return list
	}
	A := list
	B := list.Next
	invTail := B.Reverse() // = C (letztes Element der Liste)
	// Liste wird rekursiv runtergebrochen, bis A,B und C auf das gleiche Element zeigen -> Liste ist schon reversed

	A.Next = B.Next //im runtergebrochenen Fall A.Next=nil
	B.Next = A      // im runtergebrochenen Fall ist B.Next A

	return invTail
}
