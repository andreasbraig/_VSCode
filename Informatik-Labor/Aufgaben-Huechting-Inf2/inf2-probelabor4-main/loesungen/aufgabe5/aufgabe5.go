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
	if n.IsEmpty() {
		return result
	}
	if n.IsLeaf() {
		return []string{""}
	}
	for _, path := range n.Left.PathStrings() {
		result = append(result, "l"+path)
	}
	for _, path := range n.Right.PathStrings() {
		result = append(result, "r"+path)
	}
	return result
}
