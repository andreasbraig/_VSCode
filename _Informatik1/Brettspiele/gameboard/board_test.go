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

func ExampleBoard_AnyColumnContainsOnly() {
	board := Board{
		{" ", "X", " "},
		{"X", "X", "X"},
		{" ", "X", " "},
	}

	fmt.Println(board.AnyColumnContainsOnly("X"))
	fmt.Println(board.AnyColumnContainsOnly("O"))

	// Output:
	// true
	// false
}

func ExampleBoard_String() {

	b1 := Board{
		{" ", "O", " "},
		{"X", "O", " "},
		{"X", "O", " "},
	}
	b2 := Board{
		{" ", "O", " ","U"},
		{"X", "O", " ","U"},
		{"X", "O", " ","U"},
	}

	fmt.Println(b1)
	fmt.Println(b2)

	//Output:
	// +---+---+---+
	// |   | O |   |
	// +---+---+---+
	// | X | O |   |
	// +---+---+---+
	// | X | O |   |
	// +---+---+---+
	//
	// +---+---+---+---+
	// |   | O |   | U |
	// +---+---+---+---+
	// | X | O |   | U |
	// +---+---+---+---+
	// | X | O |   | U |
	// +---+---+---+---+
}
