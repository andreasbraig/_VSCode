package gamelogic

import (
	"TicTacToe/input"
	"TicTacToe/output"
	"TicTacToe/spielfeld"
	"TicTacToe/uts"
	"fmt"
)

// Hier passiert TicTacToe
func TicTacToe() {

	// Definiert ein spielfeld

	board := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	symbol := "Q"

	win := false  //Zum spielgewinn hat keiner gewonnen
	Duce := false // zum spielbeginn kann kein Unentschieden vorliegen
	run := true   // Zum spielbeginn läuft das spiel (Dies ist die oberste abbruchbedingung des Programms)
	counter := 0  //Der Rundenzähler wird auf 0 gesetzt.

	//Vorbereitung des Spielfeldes

	spielfeld.PrepareBoard(board)

	// Schleife, die das geschehen des spiels durchführt, bis das Spiel händisch beendet werden soll.

	for run {

		for !win && !Duce {

			symbol = PlayerTurn(counter) // Legt fest, welcher spieler grade am zug ist.

			board = GameMove(symbol, board) // Führt den spielzug durch

			win = spielfeld.PlayerWins(board, symbol) // Überprüft ob ein gewinn vorliegt

			Duce = spielfeld.Duce(board) // Prüft ob ein Unentschieden vorliegt.

			counter++

		}

		output.PrintCorrectly(board) //Stellt das aktuelle Spielfeld ein letztes mal dar

		if win {
			win = false
			counter = 0

			output.DisplayWin(symbol) // Zeigt den Gewinn Bildschirm

			run = uts.AskforNewRound() // Erfragt Revance

			if run {

				spielfeld.PrepareBoard(board)

			}

		} else if Duce {
			Duce = false
			counter = 0

			output.DisplayDuce() // Zeigt den Unentschieden Bildschirm

			run = uts.AskforNewRound() // Erfragt Revance

			if run {

				spielfeld.PrepareBoard(board)

			}
		}

	}

}

// Erwartet das aktuelle spielfeld
// führt einen spielzug durch und übergibt die neue Liste und das eingefügte Symbol

func GameMove(symbol string, board [][]string) [][]string {

	RowRune, ColNum := input.AskPlayer()

	err, idx, idy := input.GetExactField(RowRune, ColNum, len(board))

	if !err { // Solange kein falsches Feld oder buchstabe angebeben wurde

		if !input.CheckIsInUse(idx, idy, board) { //Solange der input nicht belegt ist

			board[idx][idy] = symbol       // Eintrag in das spielfeld
			output.DisplayGameboard(board) // Spielfeld darstellen

			return board
		}
	}

	output.DisplayGameboard(board)
	fmt.Println("Bitte korrekten wert angeben.")

	return board

}

//Erwartet die aktuelle Rundenzahl
//Übergibt das symbol des spielers, der dran ist.

func PlayerTurn(n int) string {
	if n%2 == 0 {
		return "X"
	} else {
		return "O"
	}

}
