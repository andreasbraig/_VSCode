package dllist_int_solution

import "testing"

func TestDLList_Swap(t *testing.T) {
	l1 := NewDLList()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(4)
	l1.PushBack(5)

	l1.Swap(2, 4)

	if l1.String() != "[1 4 3 2 5]" {
		t.Error("List should be [1 4 3 2 5], but is", l1.String())
	}
}
