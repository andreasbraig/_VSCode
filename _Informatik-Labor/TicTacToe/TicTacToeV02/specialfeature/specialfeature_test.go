package specialfeature

import (
	"TicTacToe/output"
	"fmt"
)

func ExampleNewBoard() {

	Board := NewBoard(4)

	fmt.Println(Board)
	output.PrintCorrectly(Board)

	//Output:
	//

}
