package dllist_int

import "testing"

func TestDLLNode_Insert_nextpointers(t *testing.T) {
	e := NewDLLNode() // e
	last := e.Insert(1) // e -> 1
	e.Insert(3)
	e.Insert(2)
	last.Insert(38)

	if e.Next.Value != 2 {
		t.Error("e.Next.Value should be 2, but is", e.Next.Value)
	}
	if e.Next.Next.Value != 3 {
		t.Error("e.Next.Next.Value should be 3, but is", e.Next.Next.Value)
	}
	if e.Next.Next.Next.Value != 1 {
		t.Error("e.Next.Next.Next.Value should be 1, but is", e.Next.Next.Next.Value)
	}
	if e.Next.Next.Next.Next.Value != 38 {
		t.Error("e.Next.Next.Next.Next.Value should be 38, but is", e.Next.Next.Next.Next.Value)
	}
	if e.Next.Next.Next.Next.Next != e {
		t.Error("e.Next.Next.Next.Next.Next should be e, but is", e.Next.Next.Next.Next.Next)
	}
}

func TestDLLNode_Insert_prevpointers(t *testing.T) {
	e := NewDLLNode()
	last := e.Insert(1)
	e.Insert(3)
	e.Insert(2)
	last.Insert(38)

	if e.Prev.Value != 38 {
		t.Error("e.Prev.Value should be 38, but is", e.Prev.Value)
	}
	if e.Prev.Prev.Value != 1 {
		t.Error("e.Prev.Prev.Value should be 1, but is", e.Prev.Prev.Value)
	}
	if e.Prev.Prev.Prev.Value != 3 {
		t.Error("e.Prev.Prev.Prev.Value should be 3, but is", e.Prev.Prev.Prev.Value)
	}
	if e.Prev.Prev.Prev.Prev.Value != 2 {
		t.Error("e.Prev.Prev.Prev.Prev.Value should be 2, but is", e.Prev.Prev.Prev.Prev.Value)
	}
	if e.Prev.Prev.Prev.Prev.Prev != e {
		t.Error("e.Prev.Prev.Prev.Prev.Prev should be e, but is", e.Prev.Prev.Prev.Prev.Prev)
	}
}
