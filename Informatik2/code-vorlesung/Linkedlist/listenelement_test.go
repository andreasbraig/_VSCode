package linkedlist

import "testing"

func TestDLElement_IsEmpty(t *testing.T) {
	e := NewDLElement()
	if !e.IsEmpty() {
		t.Error("e.IsEmpty() on new element Should return true")
	}
}

func TestDLElement_Insert_nonEmptyList(t *testing.T) {
	e := NewDLElement()
	e.Insert(1)
	if e.IsEmpty() {
		t.Error("IsEmpty() on new Element should return false")
	}
}

func TestDLElement_Insert(t *testing.T) {
	e := NewDLElement()
	last := e.Insert(1)
	e.Insert(1)
	e.Insert(3)
	e.Insert(2)
	last.Insert(38)

	if e.Next.Value != 2 {
		t.Error("e.Next.Value should be 2 but is", e.Next.Value)
	}
	if e.Next.Next.Value != 3 {
		t.Error("e.Next.Next.Value should be 3 but is", e.Next.Next.Value)
	}
	if e.Next.Next.Next.Value != 1 {
		t.Error("e.Next.Next.Next.Value should be 1 but is", e.Next.Next.Next.Value)
	}
	if e.Next.Next.Next.Next.Next.Value != 38 {
		t.Error("e.Next.Next.Next.Next.Value Should be 38 but is", e.Next.Next.Next.Next.Value)
	}
	if e.Next.Next.Next.Next.Next.Next != e {
		t.Error("e.Next.Next.Next.Next.next should be e but is", e.Next.Next.Next.Next.Next)
	}

	//TODO Check Func es geht nicht.

}
