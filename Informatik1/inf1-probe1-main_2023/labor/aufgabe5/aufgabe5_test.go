package aufgabe5

import "fmt"

func ExampleArraySums() {
	fmt.Println(ArraySums([]int{1, 3, 5, 7}))  // Soll 1 4 9 16 ausgeben.
	fmt.Println(ArraySums([]int{1, 1, 2, 80})) // Soll 1 2 4 84 ausgeben.
	fmt.Println(ArraySums([]int{7, 3, 1, 2}))  // Soll 7 10 11 13 ausgeben.
	fmt.Println(ArraySums([]int{3, 3, 0, 2}))  // Soll 3 6 6 8 ausgeben.

	// Output:
	// [1 4 9 16]
	// [1 2 4 84]
	// [7 10 11 13]
	// [3 6 6 8]
}
