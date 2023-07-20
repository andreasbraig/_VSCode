package aufgabe3

/* ERREICHBARE PUNKTE: 10 */

// In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem Graphen.
// Vervollständigen Sie die Methode FirstNeighbourStartingWith.

/***AUFGABE***/
// FirstNeighbourStartingWith soll ersten Knoten unter den Nachbarn von n liefern,
// dessen Label mit dem Präfix prefix beginnt.
// Falls keine solchen Nachbarn existieren, soll nil zurückgegeben werden.
func (n *Node) FirstNeighbourStartingWith(prefix string) *Node {
	var result *Node
	for _, e := range n.neighbours {
		if len(e.Label) >= len(prefix) { // nötig für den Vergleich
			if e.Label[:len(prefix)] == prefix {
				return e
			}
		}
	}
	return result
}
