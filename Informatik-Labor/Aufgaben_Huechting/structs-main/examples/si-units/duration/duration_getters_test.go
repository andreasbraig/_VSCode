package duration

import "fmt"

// ExampleDuration_getters erzeugt eine Duration und
// lässt sie sich in verschiedene Formate konvertieren.
func ExampleDuration_getters() {
	d := FromMS(97238765)

	fmt.Println(d.ToMS())
	fmt.Println(d.ToS())
	fmt.Println(d.ToM())
	fmt.Println(d.ToH())
	fmt.Println(d.ToD())

	// Output:
	// 97238765
	// 97238
	// 1620
	// 27
	// 1
}
