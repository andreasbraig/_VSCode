package input

import "fmt"

func AskPlayer() (int, string) {

	fieldNumber := 0
	Symbol := ""

	fmt.Println("Bitte geben sie ein feld ein:")
	fmt.Scanln(&fieldNumber)

	fmt.Println("Bitte geben sie ihr symbol ein:")
	fmt.Scanln(&Symbol)

	if CheckSymbolInput(Symbol) {

		return fieldNumber, Symbol
	}

	return -1, Symbol

}

//Erwartet einen User input der Konsole
//Überprüft ob der Input für TicTacToe Geeignet ist (Ob es ein "X" oder ein "O" ist )

func CheckSymbolInput(symbol string) bool {

	return symbol == "X" || symbol == "O"

}

func CheckFieldInput(fieldNum, fieldLength int) bool {

	if fieldNum > fieldLength*fieldLength || fieldNum <= 0 {
		return false
	}

	return true

}

//Liefert false solange ein feld noch ungenutzt ist

func CheckIsInUse(idx, idy int, board [][]string) bool {

	if board[idx][idy] == "X" || board[idx][idy] == "O" {

		return true

	}

	return false
}

// Erwartet eine Feldnummer und eine breite des spielfelds
// Liefert die Reihe in welcher der spieler sein Symbol platzieren will

func GetRow(fieldNum, fieldLength int) (bool, int) {

	rowNum := -1

	if CheckFieldInput(fieldNum, fieldLength) {

		rowNum = (fieldNum - 1) / fieldLength
		return true, rowNum
	}

	return false, -1

}

// Erwartet die
func GetExactField(fieldNum, fieldLength int) (bool, int, int) {

	err, row := GetRow(fieldNum, fieldLength)

	if !err {
		return err, -1, -1
	}

	field := (fieldNum - 1) % fieldLength

	return err, row, field

}
