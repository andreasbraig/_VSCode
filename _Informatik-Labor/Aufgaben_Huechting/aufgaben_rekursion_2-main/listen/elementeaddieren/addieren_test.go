package elementeaddieren

import "fmt"

func ExampleAddElements() {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{2, 4, 6, 8, 10}
	l3 := []int{}

	fmt.Println(AddElements(l1))
	fmt.Println(AddElements(l2))
	fmt.Println(AddElements(l3))

	// Output:
	// 15
	// 30
	// 0
}
