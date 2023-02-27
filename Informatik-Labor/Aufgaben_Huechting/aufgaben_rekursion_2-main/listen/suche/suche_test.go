package suche

import "fmt"

func ExampleSearch() {
	l1 := []int{1, 3, 4, 6, 10, 21, 38}

	fmt.Println(Search(l1, 10))
	fmt.Println(Search(l1, 50))

	// Output:
	// 4
	// 7
}
