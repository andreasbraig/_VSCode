package main

import "fmt"

func main() {

	board := [][]string{
		{"X", "O", "X"},
		{"X", "O", "O"},
		{"X", "O", "X"},
	}

	fmt.Println(board[0][1])

	fmt.Println(len(board))
}
