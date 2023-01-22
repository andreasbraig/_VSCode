package gamelogic

import (
	"TicTacToe/input"
	"TicTacToe/output"
	"TicTacToe/spielfeld"
	"TicTacToe/uts"
	"fmt"
)

func TicTacToe() {

	//Einführung der Variablen

	// Definiert ein spielfeld
	Demoboard := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	board := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	symbol := "Q"

	win := false
	end := true

	fmt.Println("Dies ist die Nummerierung der spielfelder:")
	output.PrintCorrectly(Demoboard)

	// Leert das spielfeld
	board = spielfeld.ClearBoard(board)

	// Schleife, die das geschehen des spiels durchführt, bis gewinn vorliegt

	for end {

		for !win {

			symbol, board = GameMove(board) // Führt den spielzug durch

			win = spielfeld.PlayerWins(board, symbol) // Überprüft ob ein gewinn vorliegt

		}

		output.PrintCorrectly(board) //Stellt das aktuelle Spielfeld dar

		if win {
			win = false 

			uts.DisplayWin(symbol) // Zeigt den Gewinn Bildschirm

			end = uts.AskforNewRound() // Erfragt Revance

			if end {

				board = spielfeld.ClearBoard(board)

				fmt.Println("Dies ist die Nummerierung der spielfelder:")
				output.PrintCorrectly(Demoboard)

			}

		}

	}

}

// Erwartet das aktuelle spielfeld
// führt einen spielzug durch und übergibt die neue Liste und das eingefügte Symbol

func GameMove(board [][]string) (string, [][]string) {

	fieldNum, symbol := input.AskPlayer()

	err, idx, idy := input.GetExactField(fieldNum, len(board))

	if err { // Solange kein falsches Feld oder buchstabe angebeben wurde 

		if !input.CheckIsInUse(idx,idy,board){ //Solange der input nicht belegt ist 

			board[idx][idy] = symbol // Eintrag in das spielfeld
			uts.DisplayGameboard(board) // Spielfeld darstellen

			return symbol, board
		}
	}

	uts.DisplayGameboard(board)
	fmt.Println("Bitte korrekten wert angeben.") 

	return symbol, board

}
