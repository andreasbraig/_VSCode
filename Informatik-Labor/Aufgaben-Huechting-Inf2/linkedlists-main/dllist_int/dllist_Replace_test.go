package dllist_int

import "testing"

func TestDlllist_Replace(t *testing.T) {
	l1 := NewDLList()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(4)
	l1.PushBack(5)

	l1.Replace(2, 4)

	if l1.String() != "[1 2 4 4 5]" {
		t.Error("List should be [1 2 4 4 5], but is", l1.String())
	}
}

func TestDlllist_Replace2(t *testing.T) {
	l1 := NewDLList()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(4)
	l1.PushBack(5)

	l1.Replace(1, 4)

	if l1.String() != "[1 4 3 4 5]" {
		t.Error("List should be [4 2 3 4 5], but is", l1.String())
	}

}
