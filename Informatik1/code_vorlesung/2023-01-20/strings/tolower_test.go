package strings

import "fmt"

func ExampleToLower() {

	fmt.Println(ToLower("Test"))
	fmt.Println(ToLower("TEST"))
	fmt.Println(ToLower("test"))
	fmt.Println(ToLower(""))
	fmt.Println(ToLower("teSt"))
	fmt.Println(ToLower("te5St"))
	fmt.Println(ToLower("te5SÄt"))

	

	// Output:
	// test
	// test
	// test
	// 
	// test
	// te5st
	// te5sät


}