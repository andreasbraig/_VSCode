package dllist_int

import "testing"

func TestDLLNode_IsEmpty(t *testing.T) {
	e := NewDLLNode()
	if !e.IsEmpty() {
		t.Error("IsEmpty() on new element should return true")
	}
}

func TestDLLNode_Insert_nonEmptyList(t *testing.T) {
	e := NewDLLNode()
	e.Insert(1)
	if e.IsEmpty() {
		t.Error("IsEmpty() on nonempty list should return false")
	}
}
