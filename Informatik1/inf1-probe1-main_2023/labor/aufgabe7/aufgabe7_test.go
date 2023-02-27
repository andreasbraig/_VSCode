package aufgabe7

import "fmt"

func ExampleIntersect() {
	l1 := []int{1, 3, 5, 7, 9, 11, 13}
	l2 := []int{1, 4, 7, 10, 13}

	fmt.Println(Intersect(l1, l2))

	// Output:
	// [1,7,13]
}
