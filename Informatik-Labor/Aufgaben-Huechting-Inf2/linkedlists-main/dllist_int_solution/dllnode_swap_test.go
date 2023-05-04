package dllist_int_solution

import "testing"

func TestDLLNode_Swap(t *testing.T) {
	e1 := NewDLLNode()
	e2 := e1.Insert(1)
	e3 := e2.Insert(2)
	e4 := e3.Insert(3)
	e5 := e4.Insert(4)

	swapNodes(e2, e4)

	if e1.Next != e4 {
		t.Errorf("e1.Next should be e4, but is e%d.", e1.Next.Value)
	}
	if e4.Next != e3 {
		t.Errorf("e4.Next should be e3, but is e%d.", e4.Next.Value)
	}
	if e3.Next != e2 {
		t.Errorf("e3.Next should be e2, but is e%d.", e3.Next.Value)
	}
	if e2.Next != e5 {
		t.Errorf("e2.Next should be e5, but is e%d.", e2.Next.Value)
	}
	if e5.Next != e1 {
		t.Errorf("e5.Next should be e1, but is e%d.", e5.Next.Value)
	}
}
