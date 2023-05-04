package sortieren

import "fmt"

func ExampleQuicksort() {
	l1 := []int{42, 25, 12, 32, 36, 77, 66, 102, 38}

	fmt.Println(QuickSort(l1))

	//Output:
	//[12 25 32 36 38 42 66 77 102]

}
