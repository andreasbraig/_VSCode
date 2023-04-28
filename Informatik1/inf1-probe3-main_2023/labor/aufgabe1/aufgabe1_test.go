package aufgabe1

import (
	"fmt"
)

func ExampleContainsPrefix() {
	l1 := []string{"abc", "ab", "de", "abcde", "defabcedfg", "kj"}

	fmt.Println(ContainsPrefix(l1, "abc"))
	fmt.Println(ContainsPrefix(l1, "a"))
	fmt.Println(ContainsPrefix(l1, "de"))
	fmt.Println(ContainsPrefix(l1, "z"))

	// Output:
	// true
	// true
	// true
	// false
}
