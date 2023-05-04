package dllist_int_solution

import "fmt"

// DLLNode represents an element in a doubly linked list.
type DLLNode struct {
	Prev  *DLLNode
	Value int
	Next  *DLLNode
}

// NewDLLNode creates a new empty doubly linked list element.
func NewDLLNode() *DLLNode {
	e := &DLLNode{}
	e.Prev = e
	e.Next = e
	return e
}

// IsEmpty returns true if the element is empty.
func (e *DLLNode) IsEmpty() bool {
	return e.Prev == e && e.Next == e
}

// Insert inserts a new element after the current element.
func (e *DLLNode) Insert(value int) *DLLNode {
	newElement := &DLLNode{Prev: e, Value: value, Next: e.Next}
	e.Next.Prev = newElement
	e.Next = newElement
	return newElement
}

// Remove removes the current element from the list.
// If the list is empty, nothing happens.
func (e *DLLNode) Remove() {
	if e.IsEmpty() {
		return
	}
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
}

// String returns a string representation of the element.
func (e *DLLNode) String() string {
	return fmt.Sprintf("%d", e.Value)
}

// swapNodes expects two element pointers and swaps these elements by rewriting their Pointers.
func swapNodes(e1, e2 *DLLNode) {
	// Before: A <-> e1 <-> B <-> ... <-> C <-> e2 <-> D
	// After:  A <-> e2 <-> B <-> ... <-> C <-> e1 <-> D

	A := e1.Prev
	B := e1.Next
	C := e2.Prev
	D := e2.Next

	A.Next = e2
	e2.Prev = A
	e2.Next = B
	B.Prev = e2
	C.Next = e1
	e1.Prev = C
	e1.Next = D
	D.Prev = e1
}
