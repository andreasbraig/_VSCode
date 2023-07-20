package aufgabe2

/* ERREICHBARE PUNKTE: 15 */

// In der Datei linkedlist.go finden Sie eine Implementierung einer einfach verketteten Liste.
// Zusätzlich zu den üblichen Feldern haben diese Listenelemente auch ihre Länge gespeichert.
// Implementieren Sie die Methode RemoveAt.
// Dabei muss neben dem Entfernen des Elements auch die Länge der Liste angepasst werden.

/***AUFGABE***/
// RemoveAt fügt ein neues Element mit dem Wert val an der Stelle index in die Liste ein.
// Wenn index eine ungültige Position ist, soll die Liste unverändert bleiben.
func (n *Node) RemoveAt(index int) {
	if index < 0 || n.IsEmpty() {
		return
	}
	if index == 0 {
		n.Value = n.Next.Value
		n.Next = n.Next.Next
		n.Length--
		return
	}
	if index == 1 {
		n.Next = n.Next.Next
		n.Length--
		return
	}
	nextlen := n.Next.Length
	n.Next.RemoveAt(index - 1)
	if nextlen != n.Next.Length {
		n.Length--
	}
}

// Hinweis: Das Entfernen des ersten Elements ist ein Sonderfall bzw. komplizierter.
//          Deshalb wird dieser Fall separat in einer eigenen Testfunktion getestet.
//          Falls Sie Schwierigkeiten haben, können Sie den Sonderfall also
//          auch weglassen und nur die allgemeine Lösung implementieren.
