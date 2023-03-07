package aufgabe7

import "fmt"

func ExampleDifference() {
	l1 := []int{1, 3, 5, 7, 9, 11, 13}
	l2 := []int{1, 4, 7, 10, 13}

	fmt.Println(Difference(l1, l2))
	fmt.Println(Difference(l2, l1))

	// Output:
	// [3 5 9 11]
	// [4 10]
}
