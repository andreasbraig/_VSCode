package aufgabe4

import "fmt"

func ExampleGreater() {

	fmt.Println(Greater(3, 4))
	fmt.Println(Greater(4, 3))
	fmt.Println(Greater(0, 0))
	fmt.Println(Greater(1, 0))
	// Output:
	// false
	// true
	// false
	// true
}
