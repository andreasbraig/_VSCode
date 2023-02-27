package aufgabe3

import "fmt"

func ExampleZip() {

	fmt.Println(Zip("abc", "def"))
	fmt.Println(Zip("bab", "abacc"))
	fmt.Println(Zip("", ""))

	// Output:
	// adbecf
	// baabbacc
	//
}
