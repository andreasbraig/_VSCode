package spielfeld

import "fmt"

func ExampleRowFull() {

	board := [][]string{
		{" ", "X", " "},
		{"X", "X", "X"},
		{"O", "O", "O"},
	}

	fmt.Println(RowFull(board, 1, "X"))

	//Output:
	//true

}
