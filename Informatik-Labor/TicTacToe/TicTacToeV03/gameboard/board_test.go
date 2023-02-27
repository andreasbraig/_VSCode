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

func ExampleCheckCross() {

	board := Board{
		{" ", "X", "O"},
		{"X", "O", "X"},
		{"O", "X", " "},
	}

	board1 := Board{
		{"O", "X", " "},
		{" ", "O", " "},
		{" ", "X", "O"},
	}

	fmt.Println(board.CheckCross("X"))
	fmt.Println(board.CheckCross("O"))
	fmt.Println(board1.CheckCross("X"))
	fmt.Println(board1.CheckCross("O"))

	//Output:
	//false
	//true
	//false
	//true

}

func ExampleRightCross() {
	board := Board{
		{" ", "X", "O"},
		{"X", "O", "X"},
		{"O", "X", " "},
	}

	fmt.Println(board.RightCross())

	//Output:
	//| O | O | O |

}

func ExampleBoard_String() {

	b1 := Board{
		{" ", "O", " "},
		{"X", "O", " "},
		{"X", "O", " "},
	}
	b2 := Board{
		{" ", "O", " ", "U"},
		{"X", "O", " ", "U"},
		{"X", "O", " ", "U"},
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

func ExampleDuce() {

	b1 := Board{
		{"O", "O", "O"},
		{"X", "O", "X"},
		{"X", "O", "X"},
	}
	b2 := Board{
		{"O", "O", "X", "O"},
		{"X", "O", "O", "X"},
		{"X", "O", "X", "O"},
	}

	b3 := Board{
		{"O", "O", "X", "U"},
		{"X", "O", "O", "X"},
		{"X", "O", "X", "O"},
	}

	fmt.Println(b1.Duce())
	fmt.Println(b2.Duce())
	fmt.Println(b3.Duce())

	//Output:
	//true
	//true
	//false

}

func ExamplePrepareBoard() {

	b1 := Board{
		{"O", "O", "O"},
		{"X", "O", "X"},
		{"X", "O", "X"},
	}

	b1.PrepareBoard()

	//Output:

}
