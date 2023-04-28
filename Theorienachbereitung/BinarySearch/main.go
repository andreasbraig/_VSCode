package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinarySearchtree struct {
	Root *Node
}

func main() {

	//Vergabe der Daten in den Nodes
	rootNode := Node{Data: 4}
	lNode := Node{Data: 2}
	rNode := Node{Data: 7}
	lrNode := Node{Data: 3}

	//Verkettung der Nodes zu einem Baum
	rootNode.Left = &lNode
	lNode.Right = &lrNode
	rootNode.Right = &rNode

	//Inizialisierung des Suchbaumes
	var bs BinarySearchtree
	bs.Root = &rootNode

	bs.Print()

}

func (bs BinarySearchtree) Print() {

	if bs.Root != nil {
		printSubtree(bs.Root, 0)
	}

}

func printSubtree(node *Node, level int) {
	if node == nil {
		return
	}

	printSubtree(node.Right, level+1)

	for i := 0; i < level; i++ {
		fmt.Print("   ")
	}
	fmt.Println(node.Data)

	printSubtree(node.Left, level+1)

}

func (bs BinarySearchtree) Contains(element int) bool {

	node := bs.Root

	if node == nil {
		return false
	}

	//Aufbruch des Baumes ohne die root. 
	return node.NodeContains(element)

}

func (node *Node) NodeContains(element int) bool {
	//Abfrage ob die Node leer ist 
	if node == nil {
		return false
	}
	//Ist die aktuelle Node schon unser gesuchtes Ergebnis?
	if node.Data == element {
		return true
	}

	//Bewegung in die richtung die wir erwarten würden
	if node.Data < element {
		return node.Right.NodeContains(element)
	} else {
		return node.Left.NodeContains(element)
	}

}
