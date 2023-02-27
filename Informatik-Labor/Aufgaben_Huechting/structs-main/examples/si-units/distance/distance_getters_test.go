package distance

import "fmt"

// ExampleDistance_getters erzeugt eine Distance und
// lässt sie sich in verschiedene Formate konvertieren.
func ExampleDistance_getters() {
	d := FromMM(15238765)

	fmt.Println(d.ToMM())
	fmt.Println(d.ToCM())
	fmt.Println(d.ToM())
	fmt.Println(d.ToKM())
	fmt.Println(d.ToMi())

	// Output:
	// 15238765
	// 1523876
	// 15238
	// 15
	// 9
}
