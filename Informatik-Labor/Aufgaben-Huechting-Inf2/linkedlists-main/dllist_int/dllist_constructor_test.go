package dllist_int

import "testing"

func TestDLList_IsEmpty(t *testing.T) {
	l := NewDLList()
	if !l.IsEmpty() {
		t.Error("IsEmpty() on new list should return true")
	}
}

func TestDLList_PushBack_nonEmptyList(t *testing.T) {
	l := NewDLList()
	l.PushBack(1)
	if l.IsEmpty() {
		t.Error("IsEmpty() on nonempty list should return false")
	}
}

func TestDLList_PushFront_nonEmptyList(t *testing.T) {
	l := NewDLList()
	l.PushFront(1)
	if l.IsEmpty() {
		t.Error("IsEmpty() on nonempty list should return false")
	}
}
