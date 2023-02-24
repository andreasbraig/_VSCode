package main

import (
	"fmt"
	"strings"
)

type Row []string

type Board []Row

func main() {

	letterlist := Row{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

	ChessBoard := Board{}

	fmt.Println(ChessBoard.NameChessBoard(letterlist))

}

func (ChessBoard Board) NameChessBoard(letterlist Row) Board {

	for idx := range letterlist {

		ChessBoard = append(ChessBoard, []string{}) //Die Leere Zeile zum Füllen wird erstellt

		for _, element := range letterlist { //For Schleife für Zeile

			field := element + fmt.Sprintf("%d", idx+1) //Sprintf konvertiert den Int wert von I+1 zu einem String und addiert ihn auf element, damit sie gemeinsam field ergeben

			ChessBoard[idx] = append(ChessBoard[idx], field) //Die Leere Zeile wird gefüllt

		}

	}

	return ChessBoard

}

func (row Row) String() string {

	result := "|"

	for _, element := range row {
		result += fmt.Sprintf(" %s |", element)
	}

	return result

}

func (board Board) String() string {

	width := len(board[0])

	divider := strings.Repeat("+----", width) + "+\n"

	result := divider

	for _, row := range board {

		result += fmt.Sprintf("%v\n", row)
		result += divider

	}

	return result

}
