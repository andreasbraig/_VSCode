package aufgabe5

import "fmt"

func ExampleArrayMax() {
	fmt.Println(ArrayMax([]int{1, 3, 5, 7}))
	fmt.Println(ArrayMax([]int{1, 1, 203, 80}))
	fmt.Println(ArrayMax([]int{7, 3, 1, 2}))
	fmt.Println(ArrayMax([]int{3, 3, 0, 2}))
	fmt.Println(ArrayMax([]int{}))

	// Output:
	// 7
	// 203
	// 7
	// 3
	// 0
}
