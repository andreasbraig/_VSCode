package aufgabe2

import "fmt"

func ExampleLinkedList_Filter_someElements_original() {
	l1 := MakeLinkedList("A", "B", "C")
	fmt.Println(l1.Filter(func(s string) bool { return s == "A" || s == "C" }))

	// Output:
	// [A C]
}

func ExampleLinkedList_Filter_allElements_original() {
	l1 := MakeLinkedList("A", "B", "C")
	fmt.Println(l1.Filter(func(s string) bool { return true }))

	l2 := MakeLinkedList()
	fmt.Println(l2.Filter(func(s string) bool { return true }))

	// Output:
	// [A B C]
	// []
}

func ExampleLinkedList_Filter_noElements_original() {
	l1 := MakeLinkedList("A", "B", "C")
	fmt.Println(l1.Filter(func(s string) bool { return s == "D" }))

	// Output:
	// []
}
