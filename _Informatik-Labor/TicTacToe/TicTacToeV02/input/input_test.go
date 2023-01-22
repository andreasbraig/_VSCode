package input

import "fmt"

func ExampleGetRow() {

	fmt.Println(GetRow('A'))
	fmt.Println(GetRow('B'))
	fmt.Println(GetRow('C'))

	//Output:
	//0
	//1
	//2

}

func ExampleSplitInput(){

	fmt.Println(SplitInput("a1"))
	fmt.Println(SplitInput("C,2"))

	//Output:
	//65 1
	//67 2


}