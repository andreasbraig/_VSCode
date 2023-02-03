package gameboard

import "fmt"

func ExampleBoard_AnyRowContainsOnly() {
	board := Board{
		{" ", "X", " "},
		{"X", "X", "X"},
		{" ", "X", " "},
	}

	fmt.Println(board.AnyRowContainsOnly("X"))
	fmt.Println(board.AnyRowContainsOnly("O"))

	// Output:
	// true
	// false
}
