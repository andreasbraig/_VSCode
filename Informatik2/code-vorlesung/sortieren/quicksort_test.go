package sortieren

import "fmt"

func Example_Quicksort() {

	list := []int{7, 4, 10, 2, 0, 5, 44, 6, 23, 8}
	Quicksort(list)

	fmt.Println(list)

	// Output:
	// [0 2 4 5 6 7 8 10 23 44]

}
