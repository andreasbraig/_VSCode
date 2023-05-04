package dllist_int_solution

import "testing"

func TestDLList_PopBack_string(t *testing.T) {
	l := NewDLList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.PopBack()

	if l.String() != "[1 2]" {
		t.Error("String() should return [1 2], but is", l.String())
	}
}

func TestDLList_PopFront_string(t *testing.T) {
	l := NewDLList()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)

	l.PopFront()

	if l.String() != "[2 1]" {
		t.Error("String() should return [2 1], but is", l.String())
	}
}
