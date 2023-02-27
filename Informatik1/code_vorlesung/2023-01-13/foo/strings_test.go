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
	// -1

}

func ExampleLookupString() {

	list1 := CreateEmptyStringSlice()
	list1 = append(list1, "Max Mustermann", "Monika Mustermann")
	list2 := CreateEmptyStringSlice()
	list2 = append(list2, "0123-45678", "9876-54321")

	fmt.Println(LookupString(list1, list2, "Max Mustermann"))
	fmt.Println(LookupString(list1, list2, "Monika Mustermann"))
	fmt.Println(LookupString(list1, list2, "test"))

	// Output:
	// 0123-45678
	// 9876-54321
	//

}
