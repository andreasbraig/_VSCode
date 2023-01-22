package input

import "fmt"

func ExampleGetRow() {

	board := [][]string{
		{"X", "X", "X"},
		{"O", "O", "O"},
		{"O", " ", "X"},
	}

	fmt.Println(GetRow(7, len(board)))
	fmt.Println(GetRow(4, len(board)))
	fmt.Println(GetRow(2, len(board)))
	fmt.Println(GetRow(6, len(board)))
	fmt.Println(GetRow(12, len(board)))

	//Output:
	//true 2
	//true 1
	//true 0
	//true 1
	//false -1

}

func ExampleGetExactField() {

	board := [][]string{
		{"X", "X", "X"},
		{"O", "O", "O"},
		{"O", " ", "X"},
	}

	fmt.Println(GetExactField(1, len(board)))
	fmt.Println(GetExactField(2, len(board)))
	fmt.Println(GetExactField(3, len(board)))
	fmt.Println(GetExactField(4, len(board)))
	fmt.Println(GetExactField(5, len(board)))
	fmt.Println(GetExactField(6, len(board)))
	fmt.Println(GetExactField(7, len(board)))
	fmt.Println(GetExactField(8, len(board)))
	fmt.Println(GetExactField(9, len(board)))
	fmt.Println(GetExactField(-1, len(board)))
	fmt.Println(GetExactField(11, len(board)))

	//Output:
	// true 0 0
	// true 0 1
	// true 0 2
	// true 1 0
	// true 1 1
	// true 1 2
	// true 2 0
	// true 2 1
	// true 2 2
	// false -1 -1
	// false -1 -1

}
