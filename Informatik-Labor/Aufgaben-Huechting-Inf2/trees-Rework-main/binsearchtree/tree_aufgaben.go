package binsearchtree

/*
Diese Datei enthält die Aufgaben für den Binärbaum.
Es geht darum, weitere Methoden für den Binärbaum zu implementieren.
*/

// RemoveValue entfernt ein Element mit dem Wert value aus dem Baum.
// Wenn das Element nicht gefunden wird, passiert nichts.
func (t *Tree) RemoveValue(value int) {
	// TODO
}

// NodeCount gibt die Anzahl der Elemente im Baum zurück.
func (t *Tree) NodeCount() int {
	// TODO
	return 0
}

// IsSearchTree gibt true zurück, wenn der Baum ein Suchbaum ist.
func (t *Tree) IsSearchTree() bool {
	// TODO
	return false
}

/** ACHTUNG: Ab hier sind die Funktionen alle in Element definiert.
 * Das ist eigentlich ein Fehler, sie sollten besser in element.go stehen.
 * Aber da ich die Aufgaben schon mehrfach angepasst habe und weitere Verwirrung
 * vermeiden möchte, lasse ich es jetzt so.
**/

// Height gibt die Höhe des Baums zurück.
func (t *Element) Height() int {
	// TODO
	return 0
}

// BalanceFactor gibt den Balancefaktor des Knotens zurück.
func (e *Element) BalanceFactor() int {
	// TODO
	return 0
}

// RotateLeft rotiert den Baum nach links.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateLeft(root *Element) *Element {
	// TODO
	return nil
}

// RotateRight rotiert den Baum nach rechts.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateRight(root *Element) *Element {
	// TODO
	return nil
}

// RotateLeftRight führt eine Doppelrotation nach links aus.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateLeftRight(root *Element) *Element {
	// TODO
	return nil
}

// RotateRightLeft führt eine Doppelrotation nach rechts aus.
// Die Funktion erwartet die Wurzel des (Teil-)Baums und gibt die neue Wurzel zurück.
func RotateRightLeft(root *Element) *Element {
	// TODO
	return nil
}
