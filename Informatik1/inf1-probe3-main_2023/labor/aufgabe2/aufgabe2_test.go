package aufgabe2

import "fmt"

func ExampleGetStringsBetweenEmptyLines() {
	l1 := []string{"abc", "", "def", "", "ghi"}
	l2 := []string{"", "def", "", "ghi"}
	l3 := []string{"def", "", "ghi"}
	l4 := []string{"abc", "def", "ghi"}

	fmt.Println(GetStringsBetweenEmptyLines(l1))
	fmt.Println(GetStringsBetweenEmptyLines(l2))
	fmt.Println(GetStringsBetweenEmptyLines(l3))
	fmt.Println(GetStringsBetweenEmptyLines(l4))

	// Output:
	// [def]
	// [def]
	// []
	// []

}
