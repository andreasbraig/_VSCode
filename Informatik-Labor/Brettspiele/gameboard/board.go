package gameboard

import (
	"fmt"
	"strings"
)

type Board []Row

// Liefert Spalte i.
func (board Board) Column(i int) Row {

	result := Row{}

	//Schleife um aus dem Spielfeld jedes Symbol an der Spalte i in eine neue
	//"Zeile" ergänzt wird um diese dann als Reihe testen zu können.
	for _, row := range board {
		result = append(result, row[i])
	}

	return result

}

//Liefert true falls eine Spalte in board mit Symbol gefüllt ist

func (board Board) AnyColumnContainsOnly(symbol string) bool { //In der Funktion gibt er jetzt den Datentyp an. (Methode)

	for i := range board[0] {

		if board.Column(i).ContainsOnly(symbol) { // Ruft die Funktion  Contains only im file row aus
			return true
		}
	}
	return false
}

// Liefert true, falls irgendeine Zeile in board mit
// symbol gefüllt ist.
func (board Board) AnyRowContainsOnly(symbol string) bool { //In der Funktion gibt er jetzt den Datentyp an. (Methode)

	for _, row := range board {

		if row.ContainsOnly(symbol) { // Ruft die Funktion  Contains only im file row aus
			return true
		}
	}
	return false
}

// Liefert eine Representation vom des Boards
func (board Board) String() string {

	width := len(board[0])

	divider := strings.Repeat("+---", width) + "+\n"

	result := divider

	for _, row := range board {

		result += fmt.Sprintf("%v\n", row)
		result += divider
	}
	return result
}
