package dllist_int

import "testing"

func TestDllist_insert(t *testing.T) {
	l1 := NewDLList()
	l1.PushBack(42)
	l1.PushBack(2)
	l1.PushBack(7)
	l1.PushBack(4)
	l1.PushBack(38)

	l1.Insert(2, 4)

	if l1.String() != "[42 2 7 4 4 38]" {
		t.Error("List should be [42 2 7 4 4 38], but is", l1.String())
	}
}

func TestDlllist_Remove(t *testing.T){
	l1 := NewDLList()
	l1.PushBack(42)
	l1.PushBack(2)
	l1.PushBack(7)
	l1.PushBack(4)
	l1.PushBack(38)

	l1.Remove(2)

	if l1.String() != "[42 2 4 38]" {
		t.Error("List should be [42 2 4 38], but is", l1.String())
	}
}