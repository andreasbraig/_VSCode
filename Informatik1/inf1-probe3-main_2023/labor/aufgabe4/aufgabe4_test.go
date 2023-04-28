package aufgabe4

import "fmt"

func ExampleEqual() {

	fmt.Println(Equal(3, 4))
	fmt.Println(Equal(4, 3))
	fmt.Println(Equal(0, 0))
	fmt.Println(Equal(1, 0))
	fmt.Println(Equal(1, 1))
	// Output:
	// false
	// false
	// true
	// false
	// true
}
