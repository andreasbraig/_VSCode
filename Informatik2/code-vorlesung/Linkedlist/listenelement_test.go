package linkedlist

import "testing"

func TestDLElement_IsEmpty(t *testing.T) {
	e:= NewDLElement()
	if !e.IsEmpty(){
		t.Error("e.IsEmpty() on new element Should return true")
	}
}