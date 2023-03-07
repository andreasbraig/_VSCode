package aufgabe1

import (
	"fmt"
)

func ExampleLongestAbc() {
	l1 := []string{"abc", "ab", "de", "abcde", "defabcedfg", "kj"}
	l2 := []string{"dabc", "ab", "de", "cabcde", "defabcedfg", "kj"}
	l3 := []string{"abc"}

	fmt.Println(LongestAbc(l1))
	fmt.Println(LongestAbc(l2))
	fmt.Println(LongestAbc(l3))

	// Output:
	// abcde
	//
	// abc
}
