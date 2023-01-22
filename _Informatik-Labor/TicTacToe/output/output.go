package output

import "fmt"

// Erwartet das spielfeld
// Stellt das spielfeld in der Konsole so dar, dass es aussieht wie TicTacToe

func PrintCorrectly(board [][]string) {

	fmt.Println()
	fmt.Println()

	for idx, element := range board {
		for idy, character := range element {
			fmt.Print(character)
			if idy < len(board[idx])-1 {

				fmt.Print(" | ")
			}
		}
		fmt.Println()

		if idx < len(board)-1 {
			for i := 0; i < len(board)*3; i++ {
				fmt.Print("-")
			}
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println()

}
