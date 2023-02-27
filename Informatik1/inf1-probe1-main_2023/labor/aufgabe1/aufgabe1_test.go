package aufgabe1

import (
	"fmt"
)

func ExampleMultFive() {
	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} // Produkt: 1*6*11 == 66
	v2 := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1} // Produkt: 12*7*2 = 168
	v3 := []int{3, 29, 4, 0, 42, 2, 38}                // Produkt: 3*2 = 6
	v4 := []int{}                                      // Leere Liste, Produkt sollte 1 sein

	fmt.Println(MultFive(v1))
	fmt.Println(MultFive(v2))
	fmt.Println(MultFive(v3))
	fmt.Println(MultFive(v4))

	// Output:
	// 66
	// 168
	// 6
	// 1
}
