package aufgabe3

import "fmt"

func ExampleBinTree_GetValue_existingNode_original() {
	t1 := NewBinTree()

	t1.Add(42)
	t1.Add(23)
	t1.Add(55)
	t1.Add(77)
	t1.Add(45)

	fmt.Println(t1.GetValue(""))
	fmt.Println(t1.GetValue("l"))
	fmt.Println(t1.GetValue("r"))
	fmt.Println(t1.GetValue("rl"))
	fmt.Println(t1.GetValue("rr"))

	// Output:
	// 42
	// 23
	// 55
	// 45
	// 77
}

func ExampleBinTree_GetValue_nonExistingNode_original() {
	t1 := NewBinTree()

	t1.Add(42)
	t1.Add(23)
	t1.Add(55)
	t1.Add(77)
	t1.Add(45)

	fmt.Println(t1.GetValue("ll"))

	t2 := NewBinTree()
	fmt.Println(t2.GetValue("r"))

	// Output:
	// -1
	// -1
}
