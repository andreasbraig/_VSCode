package binsearchtree

import (
	"testing"
)

func TestTree_Swap1(t *testing.T) {

	tree := NewTree()

	proot := tree.Insert(77)
	tree.Insert(50)
	e1 := tree.Insert(17)
	e2 := tree.Insert(42)
	tree.Insert(66)
	tree.Insert(103)
	tree.Insert(200)

	if tree.Root.Left.Left != e1 {
		t.Error("Tree.Root.Left.Left SHould be 17 but is", tree.Root.Left.Left.Value)
	}

	tree.Swap(e1, e2)

	if tree.Root != proot {
		t.Error("Tree Root Should be 77 but is", tree.Root.Value)
	}
	if tree.Root.Left.Left != e2 {
		t.Error("Tree.Root.Left.Left SHould now be 42 but is", tree.Root.Left.Left.Value)
	}

}

func TestTree_Swap2(t *testing.T) {

	tree := NewTree()

	proot := tree.Insert(77)
	e0 := tree.Insert(50)
	e1 := tree.Insert(17)
	tree.Insert(42)
	tree.Insert(66)
	tree.Insert(103)
	tree.Insert(200)

	tree.Swap(e0, e1)

	if tree.Root != proot {
		t.Error("Tree Root Should be 77 but is", tree.Root.Value)
	}
	if tree.Root.Left != e1 {
		t.Error("Tree.Left.Root Should be 17 but is", tree.Root.Left.Value)
	}
	if tree.Root.Left.Right.Value != 66 {
		t.Error("Tree.Root.Left.Right Should be 66 but is", tree.Root.Left.Right.Value)
	}

}

func TestTree_Swap3(t *testing.T) {

	tree := NewTree()

	proot := tree.Insert(77)
	tree.Insert(50)
	tree.Insert(17)
	tree.Insert(42)
	tree.Insert(66)
	e1 := tree.Insert(103)
	e2 := tree.Insert(200)

	tree.Swap(e1, e2)

	if tree.Root != proot {
		t.Error("Tree Root Should be 77 but is", tree.Root.Value)
	}
	if tree.Root.Right.Value != 200 {
		t.Error("Tree.Root.Right.Value Should be 200 but is", tree.Root.Right.Value)
	}
	if tree.Root.Right.Right.Value != 103 {
		t.Error("Tree.Root.Right.Value Should be 103 but is", tree.Root.Right.Right.Value)
	}
	if tree.Root.Right.Left.Left != nil {
		t.Error("Tree.Root.Left.Left Should be nil but is", tree.Root.Right.Left.Left.Value)
	}

}

func TestTree_Swap4(t *testing.T) {

	tree := NewTree()

	proot := tree.Insert(77)
	e1 := tree.Insert(50)
	tree.Insert(17)
	tree.Insert(42)
	e2 := tree.Insert(66)
	tree.Insert(103)
	tree.Insert(200)

	tree.Swap(e1, e2)

	if tree.Root != proot {
		t.Error("Tree Root Should be 77 but is", tree.Root.Value)
	}
	if tree.Root.Left.Value != 66 {
		t.Error("Tree.Root.Left.Value Should be 66 but is", tree.Root.Left.Value)
	}
	if tree.Root.Left.Right.Value != 50 {
		t.Error("Tree.Root.Left.Value Should be 50 but is", tree.Root.Left.Right.Value)
	}
	if tree.Root.Left.Left.Value != 17 {
		t.Error("Tree.Root.Left.Value Should be 17 but is", tree.Root.Left.Left.Value)
	}

}

func TestTree_Remove(t *testing.T) {
	tree := NewTree()
	tree.Insert(77)
	tree.Insert(50)
	tree.Insert(17)
	tree.Insert(42)
	tree.Insert(66)
	tree.Insert(103)
	tree.Insert(200)

	tree.RemoveValue(66)

	if !tree.Root.Left.Right.IsEmpty() {
		t.Error("Tree.Root.Left.Right Should be Nil but is", tree.Root.Left.Right.Value)
	}

}

func TestTree_Remove2(t *testing.T) {
	tree := NewTree()
	tree.Insert(77)
	tree.Insert(50)
	tree.Insert(17)
	tree.Insert(42)
	tree.Insert(66)
	tree.Insert(103)
	tree.Insert(200)

	tree.RemoveValue(17)

	if tree.Root.Left.Left.Value != 42 {
		t.Error("Tree.Root.Left.Left.Value Should be 42 but is", tree.Root.Left.Left.Value)
	}

}

func TestTree_Remove3(t *testing.T) {
	tree := NewTree()
	tree.Insert(77)
	tree.Insert(50)
	tree.Insert(17)
	tree.Insert(42)
	tree.Insert(66)
	tree.Insert(103)
	tree.Insert(200)

	tree.RemoveValue(50)

	if tree.Root.Left.Left.Value != 17 {
		t.Error("Tree.Root.Left.Left.Value Should be 42 but is", tree.Root.Left.Left.Value)
	}

}

func TestTree_Remove4(t *testing.T) {
	tree := NewTree()
	tree.Insert(77)
	tree.Insert(50)
	tree.Insert(17)
	tree.Insert(42)
	tree.Insert(66)
	tree.Insert(103)
	tree.Insert(200)

	tree.RemoveValue(42)

	if !tree.Root.Left.Left.Right.IsEmpty() {
		t.Error("Tree.Root.Left.Left.Right Should be nil but is", tree.Root.Left.Left.Right.Value)
	}

}

func TestTree_Remove5(t *testing.T) {
	tree := NewTree()
	tree.Insert(77)
	tree.Insert(50)
	tree.Insert(17)
	tree.Insert(42)
	tree.Insert(66)
	tree.Insert(103)
	tree.Insert(200)

	tree.RemoveValue(103)

	if tree.Root.Right.Value != 200 {
		t.Error("Tree.Root.Right.Value Should be 200 but is", tree.Root.Right.Value)
	}

}
