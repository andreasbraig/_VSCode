package dllist_int_solution

import "testing"

func TestDLList_PushBack_string(t *testing.T) {
	l := NewDLList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	if l.String() != "[1 2 3]" {
		t.Error("String() should return [1 2 3], but is", l.String())
	}
}

func TestDLList_PushFront_string(t *testing.T) {
	l := NewDLList()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)

	if l.String() != "[3 2 1]" {
		t.Error("String() should return [3 2 1], but is", l.String())
	}
}
