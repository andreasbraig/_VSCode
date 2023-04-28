package aufgabe3

import "fmt"

func ExampleCountGreater() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	l2 := []int{1, 3, 5, 7, 9}
	l3 := []int{2, 4, 6, 8, 10}
	l4 := []int{}

	fmt.Println(CountGreater(l1, 3))
	fmt.Println(CountGreater(l2, 3))
	fmt.Println(CountGreater(l3, 3))
	fmt.Println(CountGreater(l4, 3))

	// Output:
	// 5
	// 3
	// 4
	// 0
}
