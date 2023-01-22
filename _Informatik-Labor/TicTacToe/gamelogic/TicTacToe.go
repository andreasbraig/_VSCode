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
	board := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	symbol := "Q"

	win := false
	end := true

	fmt.Println("Dies ist die Nummerierung der spielfelder:")
	output.PrintCorrectly(board)

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

			uts.DisplayWin(symbol) // Zeigt den Gewinn Bildschirm

			end = uts.AskforNewRound() // Erfragt Revance

		}

	}

}

// Erwartet das aktuelle spielfeld
// führt einen spielzug durch und übergibt die neue Liste und das eingefügte Symbol

func GameMove(board [][]string) (string, [][]string) {

	fieldNum, symbol := input.AskPlayer()

	err, idx, idy := input.GetExactField(fieldNum, len(board))

	if err {

		board[idx][idy] = symbol
		uts.DisplayGameboard(board)

		return symbol, board
	}

	uts.DisplayGameboard(board)
	fmt.Println("Bitte korrekten wert angeben.")

	return symbol, board

}
