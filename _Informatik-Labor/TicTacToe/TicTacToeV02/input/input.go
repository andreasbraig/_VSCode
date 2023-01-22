package input

import (
	"fmt"
	"strings"
)

//Fragt den Player nach seinem Feld

func AskPlayer( /*board [][]string*/ ) (rune, int /*[][]string*/) {

	input := " "

	fmt.Println("Bitte geben sie eine Koordinate ein:")
	fmt.Scanln(&input)
	rune, number := SplitInput(input)

	/*if input == "specialfeature" {

		size := 0

		fmt.Println("Willkommen bei diesem Easteregg!")
		fmt.Println("geben sie eine neue Spielfeldgröße ein :) ")
		fmt.Scanln(&size)

		return 'Z', 99, // specialfeature.NewBoard(size)

	}*/

	return rune, number //, board

}

// Erwartet den Input String z.B. a1 oder C3
// Teilt diesen in die Rune von C und die zahl 3 auf und überigbt diese wieder.

func SplitInput(input string) (rune, int) {

	var r rune
	var num int = 0

	input = strings.ToUpper(input)

	for _, element := range input {
		if rune(element) >= 48 && rune(element) <= 57 {
			num += int(element) - '0'
		} else if rune(element) >= 65 && rune(element) <= 90 {
			r = rune(element)
		}
	}

	return r, num
}

//Erwartet einen User input der Konsole
//Überprüft ob der Input für TicTacToe Geeignet ist (Ob es ein "X" oder ein "O" ist )

func CheckSymbolInput(symbol string) bool {

	return symbol == "X" || symbol == "O"

}

//Überprüft die eingegebenen Koordinaten ob sie innerhalb des spielfeld liegen

func CheckFieldInput(RowRune rune, fieldNum, maxRows int) bool {

	bool1 := maxRows < GetRow(rune(RowRune)) || GetRow(rune(RowRune)) < 0

	bool2 := fieldNum > maxRows || fieldNum < 0

	return bool1 || bool2
}

//Liefert false solange ein feld noch ungenutzt ist

func CheckIsInUse(idx, idy int, board [][]string) bool {

	if board[idx][idy] == "X" || board[idx][idy] == "O" {

		return true

	}

	return false
}

// Erwartet einen Buchstabeben
// Liefert die Reihe in welcher der spieler sein Symbol platzieren will

func GetRow(rowNum rune) int {

	return int(rowNum) - 'A'

}

//Erwartet Koordinaten und die maximale länge

func GetExactField(RowRune rune, ColNum, maxRows int) (bool, int, int) {

	err := CheckFieldInput(RowRune, ColNum, maxRows)

	row := GetRow(RowRune)

	field := ColNum - 1

	return err, row, field

}
