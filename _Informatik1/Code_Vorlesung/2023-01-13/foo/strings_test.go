package foo

import "fmt"

func ExampleCreateStringSliceWithOneElement() {
	fmt.Println(CreateStringSliceWithOneElement("Hallo"))

	// Output:
	// [Hallo]
}

func ExampleFindString() {

	list1 := CreateEmptyStringSlice()
	list1 = append(list1, "Hallo", "aber", "nicht", "immer", "nur", "Welt", "!!!")

	fmt.Println(FindString(list1, "Hallo"))
	fmt.Println(FindString(list1, "immer"))
	fmt.Println(FindString(list1, "nur"))
	fmt.Println(FindString(list1, "was"))

	// Output:
	// 0
	// 3
	// 4
	// 0

}
