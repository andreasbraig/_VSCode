package binsearchtree

// Element is a node in a binary search tree.
type Element struct {
	Value int
	Left  *Element
	Right *Element
}

// NewElement creates a new empty element.
func NewElement() *Element {
	return &Element{}
}

// IsEmpty returns true if the element is empty.
func (e *Element) IsEmpty() bool {
	return e.Left == nil && e.Right == nil
}

// SetValue sets the value of the element.
// If the element is empty, two new children are created.
func (e *Element) SetValue(value int) {
	if e.IsEmpty() {
		e.Left = NewElement()
		e.Right = NewElement()
	}
	e.Value = value
}

// Insert inserts a new element into the tree.
// Returns a pointer to the new element.
func (e *Element) Insert(value int) *Element {
	if e.IsEmpty() {
		e.SetValue(value)
		return e
	}
	if value < e.Value {
		return e.Left.Insert(value)
	}
	return e.Right.Insert(value)
}

// GetNode returns the node with the given value.
// If the node is not found, nil is returned.
func (e *Element) GetNode(value int) *Element {
	if e.IsEmpty() {
		return nil
	}
	if e.Value == value {
		return e
	}
	if value < e.Value {
		return e.Left.GetNode(value)
	}
	return e.Right.GetNode(value)
}

// InOrder returns the elements of the tree in order.
func (e *Element) InOrder() []int {
	if e.IsEmpty() {
		return []int{}
	}
	return append(append(e.Left.InOrder(), e.Value), e.Right.InOrder()...)
}

// Get Parent erwatet ein Element und die Wurzel des Suchbaumes
// Liefert den Pointer auf den Elternknoten des Elements
func (e *Element) GetParent(root *Element) *Element {
	//TODO

	Value := e.Value
	curr := root

	for !curr.IsEmpty() {

		if curr.Left == e || curr.Right == e {
			return curr
		}
		if Value < curr.Value {
			curr = curr.Left
		} else {
			curr = curr.Right
		}

	}

	return nil

}


// Swap erwartet den Root des Baumes und zwei Elemente des Baumes
// Swap tauscht die beiden elemente
func (t *Tree) Swap(e1, e2 *Element) {
	// Speichern der Referenzen auf die Elternknoten
	p1 := e1.GetParent(t.Root)
	p2 := e2.GetParent(t.Root)

	// Aktualisieren der Verbindungen von p1 zu e1
	if p1 != nil {
		if p1.Left == e1 {
			p1.Left = e2
		} else {
			p1.Right = e2
		}
	}

	// Aktualisieren der Verbindungen von p2 zu e2
	if p2 != nil {
		if p2.Left == e2 {
			p2.Left = e1
		} else {
			p2.Right = e1
		}
	}

	// Austauschen der Verbindungen von e1 und e2
	e1.Left, e2.Left = e2.Left, e1.Left
	e1.Right, e2.Right = e2.Right, e1.Right
}


// IsLeaf erwartet ein Element und prüft ob dieses ein Blatt des Baumes ist 
//Übergibt true, wenn Das Element keine Kinder hat
func (e *Element) IsLeaf() bool {
	return e.Left.IsEmpty() && e.Right.IsEmpty()
}


// NodeCount erwartet ein Element und gibt die Anzahl der Knoten des Baumes darunter zurück
/*func (e *Element) nodeCount() int {
	// TODO
	if e == nil {
		return 0
	}
	if e.IsEmpty() {
		return 1
	}
	return e.Left.nodeCount() + e.Right.nodeCount()
	//Idee richtig, liefert aber noch zu viel!
}*/

func (e *Element) nodeCount() int {
	if e == nil {
		return 0
	}
	count := 0
	if !e.IsEmpty() {
		count = 1
	}
	count += e.Left.nodeCount() + e.Right.nodeCount()
	return count
}