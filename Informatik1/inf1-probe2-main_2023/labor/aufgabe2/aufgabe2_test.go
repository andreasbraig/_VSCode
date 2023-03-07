package aufgabe2

import "fmt"

func ExampleGetStringsBetween() {
	l1 := []string{"abc", "BEGIN", "def", "END", "ghi"}
	l2 := []string{"BEGIN", "def", "END", "ghi"}
	l3 := []string{"END", "def", "BEGIN", "ghi"}
	l4 := []string{"abc", "def", "ghi"}

	fmt.Println(GetStringsBetween(l1, "BEGIN", "END"))
	fmt.Println(GetStringsBetween(l2, "BEGIN", "END"))
	fmt.Println(GetStringsBetween(l3, "BEGIN", "END"))
	fmt.Println(GetStringsBetween(l4, "BEGIN", "END"))

	// Output:
	// [def]
	// [def]
	// []
	// []
}
