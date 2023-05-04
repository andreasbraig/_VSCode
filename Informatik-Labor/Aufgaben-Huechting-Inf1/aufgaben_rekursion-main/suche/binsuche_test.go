package suche

import "fmt"

func ExampleBinSearch() {
	l1 := []int{1, 3, 4, 6, 10, 21, 38}

	fmt.Println(BinSearch(l1, 10))

	// Output:
	// 4
}
