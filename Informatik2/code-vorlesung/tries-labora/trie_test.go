package trieslabora

import "testing"

func TestElement_HasValue(t *testing.T) {

	t1 := New()

	if t1.HasValue() != false {
		t.Error("t1 should have no Value")
	}

}

func TestElement_SetValue(t *testing.T) {
	t1 := New()

	t1.SetValue("Hello World")

	if t1.Value != "Hello World" {
		t.Error("t1.Value should be Hello World but is", t1.Value)
	}

}

func TestElement_AddValueForPath(t *testing.T) {
	t1 := New()

	t1.AddValueForPath("@aaaab", "consetetur")

	node := t1.left.left.left.left.right

	if !node.HasValue() {
		t.Error("Element is Empty")
	}
	if node.Value != "consetetur" {
		t.Errorf("element has value %s, but expected %s.", node.Value, "consetetur")
	}

}
