package aufgabe7

import "fmt"

func ExampleJoin() {
	l1 := []int{1, 3, 5, 7, 9, 11, 13}
	l2 := []int{1, 4, 7, 10, 13}

	fmt.Println(Join(l1, l2))
	fmt.Println(Join(l2, l1))

	// Output:
	// [1 3 5 7 9 11 13 4 10]
	// [1 4 7 10 13 3 5 9 11]
}
