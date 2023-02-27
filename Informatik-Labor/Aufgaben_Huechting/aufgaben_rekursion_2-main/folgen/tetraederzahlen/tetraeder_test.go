package tetraederzahlen

import "fmt"

func ExampleTetraeder() {
	for n := 1; n <= 10; n++ {
		fmt.Println(Tetraeder(n))
	}

	// Output:
	// 1
	// 4
	// 10
	// 20
	// 35
	// 56
	// 84
	// 120
	// 165
	// 220
}
