package aufgabe3

import "fmt"

func ExampleCountEven() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	l2 := []int{1, 3, 5, 7, 9}
	l3 := []int{2, 4, 6, 8, 10}
	l4 := []int{}

	fmt.Println(CountEven(l1))
	fmt.Println(CountEven(l2))
	fmt.Println(CountEven(l3))
	fmt.Println(CountEven(l4))

	// Output:
	// 4
	// 0
	// 5
	// 0
}
