package main

// Ein Kopoierter Code, für die Lösung des Damenproblems.

//Kan man den auch besser über "Methoden" lösen?

import (
	"fmt"
	"strings"
)

var hasQueen [8][8]bool
var inAttack [8][8]int

func main() {
	solve(0)
}

func placeQueen(i, j int) {
	hasQueen[i][j] = true
	// update attack counts
	// row
	for c := 0; c < 8; c++ {
		inAttack[i][c]++
	}
	// col
	for r := 0; r < 8; r++ {
		inAttack[r][j]++
	}
	inAttack[i][j] -= 2 // the (i,j) cell has been counted twice in the above iterations
	// diagonals
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if r == i && c == j {
				inAttack[r][c]++
				continue
			}
			if r-i == c-j || r-i == -(c-j) {
				inAttack[r][c]++
			}
		}
	}
}

func removeQueen(i, j int) {
	hasQueen[i][j] = false
	// update attack counts
	// row
	for c := 0; c < 8; c++ {
		inAttack[i][c]--
	}
	// col
	for r := 0; r < 8; r++ {
		inAttack[r][j]--
	}
	inAttack[i][j] += 2
	// diagonals
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if r == i && c == j {
				inAttack[r][c]--
				continue
			}
			if r-i == c-j || r-i == -(c-j) {
				inAttack[r][c]--
			}
		}
	}
}

var solNo int

func solve(r int) {
	if r == 7 {
		// can we place the queen somewhere
		for c, cnt := range inAttack[7] {
			if cnt == 0 {
				placeQueen(r, c)
				solNo++
				fmt.Println("Solution No.", solNo)
				pp2DBoolArr(hasQueen)
				removeQueen(r, c)
			}
		}

	} else {
		for c, cnt := range inAttack[r] {
			if cnt == 0 {
				placeQueen(r, c)
				solve(r + 1)      // solve the next row (recursive)
				removeQueen(r, c) // remove this queen (backtrack)
			}
		}
	}
}

// `pp2DBoolArr` pretty prints a 2D int array
func pp2DBoolArr(data [8][8]bool) {
	for i := 0; i < 8; i++ {
		row := make([]string, 8)
		for j := 0; j < 8; j++ {
			if data[i][j] {
				row[j] = "Q"
			} else {
				row[j] = "+"
			}
		}
		fmt.Println(strings.Join(row, " "))
	}
}
