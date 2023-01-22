package spielfeld

//Funktion muss mit einem Großbuichstaben anfangen, damit sie exportiert werden.

// Tic Tac Toe Spielfeld in der Konsole programmieren

// Überprüft ob eine Reihe vollständig mit dem Symbol gefüllt ist

func RowFull(board [][]string, rowNumber int, symbol string) bool {

	// Liefert true, falls eine Zeile in board mit
	// Symbol gefüllt ist
	for _, character := range board[rowNumber] {
		if character != symbol {
			return false
		}
	}
	return true
}

// Liefert true, falls irgendeine Zeile in Board mit
// symbol gefüllt ist

func AnyRowFull(board [][]string, symbol string) bool {
	for i := range board {
		if RowFull(board, i, symbol) {
			return true
		}
	}

	return false

}

// Liefert true, falls Eine vorgegebene spalte in Board mit dem symbol voll ist.

func ColumnFull(board [][]string, ColNumber int, symbol string) bool {

	if board[0][ColNumber] == symbol {
		for i := range board {
			if board[i][ColNumber] != symbol {
				return false
			}

		}
		return true
	}

	return false

}

// Liefert true, Falls eine Spalte in Board mit Symbol gefüllt ist.

func AnyColumnFull(board [][]string, symbol string) bool {
	for i := range board {
		if ColumnFull(board, i, symbol) {
			return true
		}
	}

	return false
}

//Liefert True, falls ein Diagonaler gewinn vorliegt.

func LeftCrossFull(board [][]string, symbol string) bool {

	for i := range board {
		if board[i][i] != symbol { // Für Kreuz von links nach rechts.
			return false
		}
	}
	return true

}

// Liefert true, falls das spielfeld ein Diagonalen gewinn hat 
func RightCrossFull(board [][]string, symbol string) bool {

	x := len(board) - 1

	for i := range board {
		if board[i][x-i] != symbol { // Für Kreuz von Rechts nach Links.
			return false
		}
	}
	return true

}

// Überpruft, ob einer der möglichen gewinnfälle eingetreten ist. 

func PlayerWins(board [][]string, symbol string) bool {

	if AnyRowFull(board, symbol) {
		return true
	} else if AnyColumnFull(board, symbol) {
		return true
	} else if LeftCrossFull(board, symbol) {
		return true
	} else if RightCrossFull(board, symbol) {
		return true
	}
	return false
}

//Erwartet das spielfeld. Leert das spielfeld vollständig 

func ClearBoard(board [][]string) [][]string {

	for ix := range board {
		for iy := range board[ix] {
			board[ix][iy] = " "
		}
	}

	return board

}
