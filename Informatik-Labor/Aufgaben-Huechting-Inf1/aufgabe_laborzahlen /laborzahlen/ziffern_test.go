package laborzahlen

import "fmt"

func ExampleDigits() {
	fmt.Println(Digits(123))
	fmt.Println(Digits(45123))

	//Output:
	// [1 2 3]
	// [4 5 1 2 3]

}

func ExampleToString() {
	fmt.Println(ToString(123))
	fmt.Println(ToString(45123))

	// Output:
	// 123
	// 45123

}
