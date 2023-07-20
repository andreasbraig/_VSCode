package aufgabe1

/* ERREICHBARE PUNKTE: 10 */

// In der Datei linkedlist.go finden Sie eine Implementierung einer einfach verketteten Liste.
// Implementieren Sie die Methode LengthGreater5.

/***AUFGABE***/
// LengthGreater5 gibt true zurück, wenn die Länge der Liste größer als 5 ist, sonst false.
func (n *Node) LengthGreater5() bool {
	
	return n.Count() > 5
}

func (n *Node) Count() int{
	if n.IsEmpty(){
		return 0
	}
	return n.Next.Count() + 1
}