package uts

import (
	"TicTacToe/output"
	"fmt"
)

// erwartet ein spielfeld
// Stellt den aktuellen stand des spieles dar

func DisplayGameboard(board [][]string) {
	fmt.Println("Dies ist das aktuelle spielfeld:")
	output.PrintCorrectly(board)
}

// Erwartet den spieler der gewonnen hat
// Stellt den Gewinn bildschirm dar

func DisplayWin(symbol string) {
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println(symbol, "hat gewonnen!")
			fmt.Println()
			fmt.Println()
			fmt.Println()
		
}

// Soll den spieler fragen, ob er noch eine runde spielen möchte. 

func AskforNewRound() bool{

	var input string

	fmt.Println("möchten Sie erneut spielen? Y/N")
			fmt.Scanln(&input)

			if input == "Y" {
				fmt.Println("Die neue Runde startet...")
				return true 
			}
			
			
			fmt.Println("Danke, dass sie TicTacToe by Andreas Braig gespielt haben!")
			return false

}