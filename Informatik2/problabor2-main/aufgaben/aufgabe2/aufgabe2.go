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

	//Rekursionsanker, die leere Liste oder die Liste mit nur einem element muss nicht gedreht werden. 
	if list.Empty() {
		return list
	}
	if list.Next.Empty(){
		return list 
	}

	//Warum 2 Rekursionsanker? 
	//Weil wir unten im Rekursionsschritt die Liste zerlegen und der hinter Teil dann noch
	//länge >= 1 haben muss, damit alles funktioniert. 

	//Die beiden ersten Elemente werden gemerkt
	A := list
	B := list.Next

	//Alles ab B Invertieren
	invTail := B.Reverse()

	// A zwischen B und das Ende hängen.
	// Das geht weil wir wissen, dass B jetzt das letzte element der invertierten ist und invTail das erste.
	A.Next = B.Next
	B.Next = A

	return invTail

}

func (list *LinkedList) ITReverse() *LinkedList {

	var prev *LinkedList
	current := list

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev

}
