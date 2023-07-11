package aufgabe1

import "fmt"

func ExampleLinkedList_FirstDuplicate_noDuplicates_original() {
	l1 := MakeLinkedList("A", "B", "C")
	fmt.Println(l1.FirstDuplicate())

	l2 := MakeLinkedList()
	fmt.Println(l2.FirstDuplicate())

	// Output:
	// 3
	// 0
}

func ExampleLinkedList_FirstDuplicate_duplicates_original() {
	l1 := MakeLinkedList("A", "B", "B", "C")

	fmt.Println(l1.FirstDuplicate())

	// Output:
	// 1
}
