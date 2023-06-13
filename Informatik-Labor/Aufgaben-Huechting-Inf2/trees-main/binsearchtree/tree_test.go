package binsearchtree

import (
	"fmt"
	"testing"
)

func TestTree_NewTree(t *testing.T) {
	tree := NewTree()
	if !tree.IsEmpty() {
		t.Error("NewTree should return an empty tree")
	}
}

func TestTree_Insert(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	if tree.IsEmpty() {
		t.Error("Tree should not be empty")
	}
	if tree.Root.Value != 10 {
		t.Error("Root should have value 10, but has", tree.Root.Value)
	}
}

func TestTree_GetMinNode(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	min := tree.Root.GetMinNode()
	if min.Value != 3 {
		t.Error("Min should be 3, but is", min.Value)
	}
}

func TestTree_GetMaxNode(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	max := tree.Root.GetMaxNode()
	if max.Value != 17 {
		t.Error("Max should be 17, but is", max.Value)
	}
}

func TestTree_GetParent(t *testing.T) {

	tree := NewTree()

	e1 := tree.Insert(77)  // 77 Ist root
	e2 := tree.Insert(50)  // Links unter 77
	e3 := tree.Insert(17)  // Links unter 50
	e4 := tree.Insert(42)  // Rechts unter 17
	e5 := tree.Insert(66)  // Rechts unter 50
	e6 := tree.Insert(103) // Rechts unter 77
	e7 := tree.Insert(200) // Rechts unter 103

	proot := e1
	p50 := e2
	p17 := e3
	p103 := e6

	fmt.Println(proot.Value)

	if e2.GetParent(tree.Root) != proot {
		t.Error("Parent of 50 Should be 77 but is:", e2.GetParent(tree.Root).Value)
	}
	if e3.GetParent(tree.Root) != p50 {
		t.Error("Parent of 17 should be 50 but is:", e3.GetParent(tree.Root).Value)
	}
	if e4.GetParent(tree.Root) != p17 {
		t.Error("Parent of 42 should be 17 but is:", e4.GetParent(tree.Root).Value)
	}
	if e5.GetParent(tree.Root) != p50 {
		t.Error("Parent of 66 should be 50 but is:", e5.GetParent(tree.Root).Value)
	}
	if e6.GetParent(tree.Root) != proot {
		t.Error("Parent of 103 should be 77 but is:", e6.GetParent(tree.Root).Value)
	}
	if e7.GetParent(tree.Root) != p103 {
		t.Error("Parent of 103 should be 77 but is:", e6.GetParent(tree.Root).Value)
	}

}

func TestTree_SwapAll(t *testing.T) {

	TestTree_Swap1(t)
	TestTree_Swap2(t)
	TestTree_Swap3(t)
	TestTree_Swap4(t)

}

func TestTree_RemoveAll(t *testing.T) {

	TestTree_Remove(t)
	TestTree_Remove2(t)
	TestTree_Remove3(t)
	TestTree_Remove4(t)

}

func TestTree_NodeCount(t *testing.T) {

	tree := NewTree()
	tree.Insert(77)
	tree.Insert(50)
	tree.Insert(17)
	tree.Insert(42)
	tree.Insert(66)
	tree.Insert(103)
	tree.Insert(200)

	if tree.NodeCount() != 7 {
		t.Error("Tree should have 7 Nodes but has", tree.NodeCount())
	}

}
