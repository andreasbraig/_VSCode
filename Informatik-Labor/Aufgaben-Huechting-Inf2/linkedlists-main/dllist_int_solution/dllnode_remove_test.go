package dllist_int_solution

import "testing"

// Creates a list with one element and removes that element.
// Checks if the list is empty afterwards.
func TestDLLNode_Remove_oneElement_element(t *testing.T) {
	e1 := NewDLLNode()
	e2 := e1.Insert(1)
	e2.Remove()
	if !e1.IsEmpty() {
		t.Error("IsEmpty() on empty list should return true")
	}
}

// Creates a list with one element and removes the dummy element.
// Checks if the other element is an empty list afterwards.
func TestDLLNode_Remove_oneElement_dummy(t *testing.T) {
	e1 := NewDLLNode()
	e2 := e1.Insert(1)
	e1.Remove()
	if !e2.IsEmpty() {
		t.Error("IsEmpty() on empty list should return true")
	}
}

// Creates an empty list and calls Remove().
// Checks if the list is still empty afterwards.
func TestDLLNode_Remove_emptyList(t *testing.T) {
	e := NewDLLNode()
	e.Remove()
	if !e.IsEmpty() {
		t.Error("IsEmpty() on empty list should return true")
	}
}
