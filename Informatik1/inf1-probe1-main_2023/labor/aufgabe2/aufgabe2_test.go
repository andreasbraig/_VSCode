package aufgabe2

import "fmt"

func ExampleColour() {
	fmt.Println(Colour(0))
	fmt.Println(Colour(9))
	fmt.Println(Colour(10))
	fmt.Println(Colour(20))
	fmt.Println(Colour(30))
	fmt.Println(Colour(40))

	// Output:
	// Rot
	// Rot
	// Rot-Gelb
	// Grün
	// Gelb
	// Rot
}
