package main

import (
	"fmt"
	"strings"
)

type Row []string

type board []Row

func main() {

	EmptyChessBoard := board{{}, {}, {}, {}, {}, {}, {}, {}}

	letterlist := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	fmt.Println(EmptyChessBoard.NameChessBoard(letterlist))

}

func (Board board) NameChessBoard(letterlist []string) board {

	for i := range letterlist {

		for _, e := range letterlist {

			field := e + fmt.Sprintf("%d", i+1)

			Board[i] = append(Board[i], field)

		}

	}

	return Board

}

func (row Row) String() string {

	result := "|"

	for _, element := range row {
		result += fmt.Sprintf(" %s |", element)
	}

	return result

}

// Liefert eine Representation vom des Boards
func (Board board) String() string {

	width := len(Board[0])

	divider := strings.Repeat("+----", width) + "+\n"

	result := divider

	for _, row := range Board {

		result += fmt.Sprintf("%v\n", row)
		result += divider
	}
	return result
}
