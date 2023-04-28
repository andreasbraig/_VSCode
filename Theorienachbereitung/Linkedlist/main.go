package main

import "fmt"

type NodeInt struct {
	data int 
	next *NodeInt
}

type LinkedList struct {
	head *NodeInt
}

func (ll LinkedList) print() {

	tmp := ll.head

	for tmp != nil {

		fmt.Print(tmp.data, "-->")

		tmp = tmp.next
	}

	fmt.Print("NIL\n")
	return
}

func (ll LinkedList) GetLength() int {

	tmp := ll.head
	count := 0

	for tmp != nil {
		count++
		tmp = tmp.next
	}

	return count

}

func (ll *LinkedList) pushBack(newData int) {

	var newNode NodeInt
	newNode.data = newData

	if ll.head == nil {
		ll.head = &newNode
		return
	}

	tmp := ll.head
	for tmp.next != nil {
		tmp = tmp.next
	}

}

func (ll *LinkedList) PushFront(newData int) {
	newNode := NodeInt{data: newData, next: ll.head}
	ll.head = &newNode
	return
}

func (ll *LinkedList) InsertAt(newData int, index int) {

	if index < 0 || index > ll.GetLength() {
		panic("Index out of Range")
	}

	if index == 0 {
		ll.PushFront(newData)
		return
	}

	newNode := &NodeInt{data: newData}

	tmp := ll.head
	for i := 0; i<index;i++{
		tmp = tmp.next
	}

	newNode.next = tmp.next

	tmp.next = newNode
}

func (ll *LinkedList) Contains(element int) bool {
	tmp := ll.head

	for tmp != nil {
		if tmp.data == element {
			return true 
		}
		tmp = tmp.next
	}

	return false
}

func (ll LinkedList) GetElement(index int) int {
	if index < 0 || index > ll.GetLength() {
		panic("Index out of Range")
	}

	tmp := ll.head
	for index > 0 {
		tmp = tmp.next
		index -- 
	}

	return tmp.data

	
}

func (ll *LinkedList) RemoveAll(element int) {

	len := ll.GetLength()

	for i := 0 ; i < len; i++ {
		if ll.GetElement(i) == element{
			ll.RemoveAt(i)
			len --
		}
	}

}

func (ll *LinkedList) RemoveAt(index int) {

	if index < 0 || index > ll.GetLength() {
		panic("Index out of Range")
	}

	if index == 0{
		ll.head = ll.head.next
		return
	}

	tmp := ll.head
	for i := 0; i < index-1; i++{
		tmp = tmp.next
	}

	tmp.next = tmp.next.next

}

func main() {
	var aNode NodeInt
	aNode.data = 1
	var bNode NodeInt
	bNode.data = 2
	aNode.next = &bNode
	var ll LinkedList
	ll.head = &aNode


	ll.print()
	fmt.Println(ll.GetLength())

	fmt.Println(ll.Contains(2))


}
