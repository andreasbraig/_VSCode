package fibonacci

import "fmt"

func ExampleFib() {
	for n := 0; n <= 10; n++ {
		fmt.Println(Fib(n))
	}

	// Output:
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// 55
	// 89
}
