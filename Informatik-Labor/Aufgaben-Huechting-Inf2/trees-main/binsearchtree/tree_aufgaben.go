package binsearchtree

/*
Diese Datei enthält die Aufgaben für den Binärbaum.
Es geht darum, weitere Methoden für den Binärbaum zu implementieren.

Beachten Sie, dass diese Datei sowohl Methoden für den Baum als auch für die Elemente enthält.
*/

// GetMinNode gibt das Element mit dem kleinsten Wert zurück.
// Wenn der Baum leer ist, wird nil zurückgegeben.
func (e *Element) GetMinNode() *Element {
	// TODO

	if e.IsEmpty() {
		return nil
	}

	//Rekursionsanker
	if e.Left.IsEmpty() {
		return e
	}

	//Rekursionsaufruf
	return e.Left.GetMinNode()
}

// GetMaxNode gibt das Element mit dem größten Wert zurück.
// Wenn der Baum leer ist, wird nil zurückgegeben.
func (e *Element) GetMaxNode() *Element {
	// TODO

	if e.IsEmpty() {
		return nil
	}

	//Rekursionsanker
	if e.Right.IsEmpty() {
		return e
	}

	//Rekursionsaufruf
	return e.Right.GetMaxNode()
}

// RemoveValue entfernt ein Element mit dem Wert value aus dem Baum.
// Wenn das Element nicht gefunden wird, passiert nichts.
func (t *Tree) RemoveValue(value int) {
	// TODO

	el := t.Root.GetNode(value)

	if t.Root.Value == value {
		//TODO
		return
	}
	if el == nil {
		return
	}
	if el.IsLeaf() {
		//remove
		el.Left = nil
		el.Right = nil
		el.Value = 0
		return
	}
	if el.Right.IsEmpty() {
		t.Swap(el.Left, el)
		t.RemoveValue(value)
	} else {
		t.Swap(el.Right, el)
		t.RemoveValue(value)
	}
}

// NodeCount gibt die Anzahl der Elemente im Baum zurück.
func (t *Tree) NodeCount() int {
	return t.Root.nodeCount()
}

// IsSearchTree gibt true zurück, wenn der Baum ein Binärbaum ist.
func (t *Tree) IsSearchTree() bool {
	// TODO
	return false
}

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
