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

// Liefert die Hauptdiagoale von links nach rechts als Reihe
func (board Board) LeftCross() Row {

	result := Row{}

	for i := range board {
		result = append(result, board[i][i])
	}

	return result

}

// Liefert die Hauptdiagonale von Rechts oben nach Links unten
func (board Board) RightCross() Row {

	result := Row{}

	x := len(board) - 1

	for i := range board {
		result = append(result, board[i][x-i])
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

// Liefert True, wenn ein Gewinn in Diagonalform vorliegt.
func (board Board) CheckCross(symbol string) bool {

	return board.LeftCross().ContainsOnly(symbol) || board.RightCross().ContainsOnly(symbol)

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

// Erwartet ein spielfeld
// Liefert true oder False, falls das Spielfeld voll ist.
func (board Board) Duce() bool {
	counter := 0

	for _, element := range board {
		for _, character := range element {

			if character == "X" || character == "O" {

				counter++

			}
		}
	}

	return counter == len(board)*len(board[0])

}

func (board Board) ClearBoard() Board {

	for ix := range board {
		for iy := range board[ix] {
			board[ix][iy] = " "
		}
	}

	return board

}

func (board Board) PrepareBoard() Board {

	Demoboard := Board{
		{" ", "1", "2", "3"},
		{"A", " ", " ", " "},
		{"B", " ", " ", " "},
		{"C", " ", " ", " "},
	}

	fmt.Println("______________________________________________________________")
	fmt.Println()
	fmt.Println("Dies ist die Nummerierung der spielfelder:")
	fmt.Println("Die eingabe der Felder erfolgt in Koordinaten: z.B. A1 oder C3")
	fmt.Println(Demoboard)

	return board.ClearBoard()
}
