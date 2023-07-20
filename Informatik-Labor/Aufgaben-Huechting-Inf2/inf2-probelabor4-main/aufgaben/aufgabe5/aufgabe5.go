package aufgabe5

/* ERREICHBARE PUNKTE: 20 */

// In der Datei bintree.go finden Sie eine Implementierung eines binären Suchbaumes.
// Vervollständigen Sie die Methode PathStrings.

/***AUFGABE***/
// PathStrings liefert eine Liste aller existierenden Pfade im Baum.
// Ein Pfad wird dabei durch einen String der Form "llr" oder "rlr" o.Ä. dargestellt.
// Dabei steht l für "links" und r für "rechts", der String beschreibt den Weg durch den Baum.
func (n *Node) PathStrings() []string {
	result := []string{}
	//SONDERFALL
	if n.IsEmpty() {
		return result // Falls der Baum leer ist 
	}
	//REKURSIONSANKER
	if n.IsLeaf() {
		return []string{""} //Wenn der Baum ein Leaf ist erstelle eine Liste mit leerem String an stelle 0 
	}
	//REKURSIONSAUFRUF LINKS!
	leftlist := n.Left.PathStrings() // Liefere eine Liste an "Leafs" links vom Baum 
	//Im Rekursionsanker ist diese liste genau ein element lang und dieser String ist ""
	for _, path := range leftlist {
		dump := "l" + path //Schreibe l VOR!! Jedes element der Liste 
		result = append(result, dump) // Hänge das ergebnis an Result an
	}
	//REKURSIONSAUFRUF RECHTS!
	rightlist := n.Right.PathStrings() // Liefere eine Liste an "Leafs" rechts vom baum 
	//Im Rekursionsanker ist diese liste genau ein element lang und dieser String ist ""
	for _, path := range rightlist {
		dump := "r" + path
		result = append(result, dump)// Hänge das ergebnis an Result an
	}

	return result //Übergebe das ergebnis 
}
