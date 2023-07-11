// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe3

/* Aufgabenstellung:
 * In der Datei bintree.go ist ein Datentyp für einen binären Suchbaum definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode GetValue().
 *
 * Diese Methode soll einen String erwarten, der nur aus den Zeichen 'l' und 'r' besteht.
 * Solche ein String beschreibt einen Pfad von der Wurzel zu einem Knoten.
 * Beispielweise beschreibt der String "rll" von der Wurzel aus den Knoten,
 * den man erreicht, wenn man einmal nach rechts und dann zweimal nach links absteigt.
 *
 * Die Methode soll diesen Knoten finden und den dortigen Wert liefern.
 * Gibt es keinen passenden Knoten, soll -1 geliefert werden.
 */

// Liefert den Wert zum gegebenen Baum-Pfad.
func (tree *BinTree) GetValue(path string) int {
	// TODO
	if tree.Empty(){
		return -1
	}
	if len(path) == 0 {
		return tree.Value
	}

	head, tail := string(path[0]), path[1:]

	if head == "r"{
		tree.Right.GetValue(tail)
	}
	if head == "l" {
		tree.Left.GetValue(tail)
	}
	return 0 

}
