package dllist_int

import "fmt"

// A doubly linked list.
// Contains a pointer to the anchor element.
type DLList struct {
	anchor *DLLNode
}

// NewDLList creates a new empty doubly linked list.
func NewDLList() *DLList {
	l := &DLList{}
	l.anchor = NewDLLNode()
	return l
}

// IsEmpty returns true if the list is empty.
func (l *DLList) IsEmpty() bool {
	return l.anchor.IsEmpty()
}

// PushBack adds a new element at the end of the list.
func (l *DLList) PushBack(value int) {
	l.anchor.Prev.Insert(value)
}

// PushFront adds a new element at the beginning of the list.
func (l *DLList) PushFront(value int) {
	// TODO

	// Hinweis: Sie können sich an der Methode PushBack() orientieren.
}

// PopBack removes the last element from the list.
// If the list is empty, nothing happens.
func (l *DLList) PopBack() {
	// TODO

	// Hinweis: Sie können sich an der Methode PushBack() orientieren.
}

// PopFront removes the first element from the list.
// If the list is empty, nothing happens.
func (l *DLList) PopFront() {
	// TODO

	// Hinweis: Sie können sich an der Methode PushBack() orientieren.
}

// Get expects an index and returns the value at that index.
// If the index is out of bounds, it returns -1.
func (l *DLList) Get(index int) int {
	// TODO

	// Hinweis: Schreiben Sie eine Schleife, beim Anker beginnend index mal
	// von Element zu Element springt.
	// Sie können Sich an der String()-Methode orientieren.
	return -1
}

// Swap expects two values, finds their corresponding elements and
// swaps these elements by rewriting their Pointers.
func (l *DLList) Swap(value1, value2 int) {
	// TODO

	// Hinweis: Verwenden Sie den gleichen Ansatz wie bei Get(), um die Elemente zu finden.
	// Verwenden Sie dann die Funktion swapDLLNodes(), um die Pointer zu tauschen.

	// Anmerkung: Es bietet sich an, eine Hilfsfunktion zu schreiben, die einen Index
	// erwartet und einen Pointer auf das entsprechende Element zurückgibt.
	// Diese Funktion können Sie dann sowohl in Get() als auch in Swap() verwenden.
}

// String returns a string representation of the list.
func (l *DLList) String() string {
	valuestrings := []string{}
	for e := l.anchor.Next; e != l.anchor; e = e.Next {
		valuestrings = append(valuestrings, e.String())
	}
	return fmt.Sprintf("%v", valuestrings)
}
