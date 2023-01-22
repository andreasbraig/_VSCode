package spielfeld

import "fmt"

func ExampleRowFull() {

	board := [][]string{
		{"X", "X", "X"},
		{"O", "O", "O"},
		{"O", " ", "X"},
	}

	c1 := "X"
	c2 := "O"

	fmt.Println(RowFull(board, 0, c1))
	fmt.Println(RowFull(board, 1, c1))
	fmt.Println(RowFull(board, 1, c2))
	fmt.Println(RowFull(board, 2, c1))

	//Output:
	//true
	//false
	//true
	//false

}

func ExampleAnyRowFull() {

	board := [][]string{
		{"X", "X", "X"},
		{"O", "O", "O"},
		{"O", "X", "X"},
	}

	c1 := "X"
	c2 := "O"
	c3 := "A"

	fmt.Println(AnyRowFull(board, c1))
	fmt.Println(AnyRowFull(board, c2))
	fmt.Println(AnyRowFull(board, c3))

	//Output:
	//true
	//true
	//false

}

func ExampleColumnFull() {
	board := [][]string{
		{"X", "O", "X"},
		{"X", "O", "O"},
		{"X", "O", "X"},
	}

	c1 := "X"
	c2 := "O"

	fmt.Println(ColumnFull(board, 0, c1))
	fmt.Println(ColumnFull(board, 0, c2))
	fmt.Println(ColumnFull(board, 1, c2))

	//Output:
	//true
	//false
	//true

}

func ExampleAnyColumnFull() {

	board := [][]string{
		{"X", "O", "X"},
		{"X", " ", "O"},
		{"X", "O", "X"},
	}

	board1 := [][]string{
		{"X", " ", "O"},
		{" ", "O", "O"},
		{"X", "X", "O"},
	}

	c1 := "X"
	c2 := "O"

	fmt.Println(AnyColumnFull(board, c1))
	fmt.Println(AnyColumnFull(board, c2))
	fmt.Println(AnyColumnFull(board1, c1))
	fmt.Println(AnyColumnFull(board1, c2))

	// Output:
	//true
	//false
	//false
	//true
}

func ExampleLeftCrossFull() {

	board := [][]string{
		{"X", " ", " "},
		{" ", "X", " "},
		{" ", " ", "X"},
	}

	board1 := [][]string{
		{"O", " ", "X"},
		{" ", "O", "X"},
		{"X", " ", "O"},
	}

	c1 := "X"
	c2 := "O"

	fmt.Println(LeftCrossFull(board, c1))
	fmt.Println(LeftCrossFull(board1, c2))
	fmt.Println(LeftCrossFull(board1, c1))

	//Output:
	//true
	//true
	//false

}

func ExampleRightCrossFull() {

	board := [][]string{
		{"X", " ", "O"},
		{" ", "O", " "},
		{"O", " ", "X"},
	}

	board1 := [][]string{
		{"O", " ", "X"},
		{" ", "X", "X"},
		{"X", " ", "O"},
	}

	c1 := "X"
	c2 := "O"

	fmt.Println(RightCrossFull(board, c1))
	fmt.Println(RightCrossFull(board, c2))
	fmt.Println(RightCrossFull(board1, c2))
	fmt.Println(RightCrossFull(board1, c1))

	//Output:
	//false
	//true
	//false
	//true

}
