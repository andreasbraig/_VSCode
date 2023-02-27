package dreieckszahlen

import "fmt"

func ExampleDreieck() {
	for n := 1; n <= 10; n++ {
		fmt.Println(Dreieck(n))
	}

	// Output:
	// 1
	// 3
	// 6
	// 10
	// 15
	// 21
	// 28
	// 36
	// 45
	// 55
}
