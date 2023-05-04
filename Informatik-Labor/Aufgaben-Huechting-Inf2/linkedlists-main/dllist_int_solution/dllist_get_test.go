package dllist_int_solution

import "testing"

func TestDLList_Get_goodcases(t *testing.T) {
	l := NewDLList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	if l.Get(0) != 1 {
		t.Error("Get(0) should return 1, but is", l.Get(0))
	}
	if l.Get(1) != 2 {
		t.Error("Get(1) should return 2, but is", l.Get(1))
	}
	if l.Get(2) != 3 {
		t.Error("Get(2) should return 3, but is", l.Get(2))
	}
}

func TestDLList_Get_outofbounds(t *testing.T) {
	l := NewDLList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	if l.Get(-1) != -1 {
		t.Error("Get(-1) should return -1, but is", l.Get(-1))
	}
	if l.Get(3) != -1 {
		t.Error("Get(3) should return -1, but is", l.Get(3))
	}
}
