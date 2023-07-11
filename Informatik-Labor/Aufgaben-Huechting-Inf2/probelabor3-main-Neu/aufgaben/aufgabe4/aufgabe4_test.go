package aufgabe4

import "fmt"

func ExampleBinTree_HasEvenElement_original() {
	t1 := NewBinTree()
	fmt.Println(t1.HasEvenElement())

	t1.Add(41)
	fmt.Println(t1.HasEvenElement())
	t1.Add(23)
	fmt.Println(t1.HasEvenElement())
	t1.Add(55)
	fmt.Println(t1.HasEvenElement())
	t1.Add(78)
	fmt.Println(t1.HasEvenElement())

	// Output:
	// false
	// false
	// false
	// false
	// true
}
