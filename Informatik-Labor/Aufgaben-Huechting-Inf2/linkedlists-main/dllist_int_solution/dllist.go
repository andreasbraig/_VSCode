package dllist_int_solution

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
	l.anchor.Insert(value)
}

// PopBack removes the last element from the list.
// If the list is empty, nothing happens.
func (l *DLList) PopBack() {
	l.anchor.Prev.Remove()
}

// PopFront removes the first element from the list.
// If the list is empty, nothing happens.
func (l *DLList) PopFront() {
	l.anchor.Next.Remove()
}

// Get expects an index and returns the value at that index.
// If the index is out of bounds, it returns -1.
func (l *DLList) Get(index int) int {
	if index < 0 {
		return -1
	}
	e := l.anchor.Next
	for index > 0 && e != l.anchor {
		e = e.Next
		index--
	}
	if e == l.anchor {
		return -1
	}
	return e.Value
}

// Swap expects two values, finds their corresponding elements and
// swaps these elements by rewriting their Pointers.
func (l *DLList) Swap(value1, value2 int) {
	e1 := l.anchor.Next
	for e1 != l.anchor && e1.Value != value1 {
		e1 = e1.Next
	}
	e2 := l.anchor.Next
	for e2 != l.anchor && e2.Value != value2 {
		e2 = e2.Next
	}
	swapNodes(e1, e2)
}

// String returns a string representation of the list.
func (l *DLList) String() string {
	valuestrings := []string{}
	for e := l.anchor.Next; e != l.anchor; e = e.Next {
		valuestrings = append(valuestrings, e.String())
	}
	return fmt.Sprintf("%v", valuestrings)
}
