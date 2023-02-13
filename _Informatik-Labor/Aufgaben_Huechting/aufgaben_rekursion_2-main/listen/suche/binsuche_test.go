package suche

import "fmt"

func ExampleBinSearch() {
	l1 := []int{1, 3, 4, 6, 10, 21, 38}

	fmt.Println(BinSearch(l1, 1))
	fmt.Println(BinSearch(l1, 3))
	fmt.Println(BinSearch(l1, 4))
	fmt.Println(BinSearch(l1, 6))
	fmt.Println(BinSearch(l1, 10))
	fmt.Println(BinSearch(l1, 21))
	fmt.Println(BinSearch(l1, 38))
	fmt.Println(BinSearch(l1, 50))
	fmt.Println(BinSearch(l1, 9))

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 7
}
