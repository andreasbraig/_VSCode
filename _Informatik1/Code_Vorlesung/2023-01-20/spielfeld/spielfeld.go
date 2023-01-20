package spielfeld

// Tic Tac Toe Spielfeld in der Konsole programmieren 

func RowFull(board [][]string,rowNumber int, symbol string) bool {

	// Liefert true, falls eine Zeile in board mit 
	// Symbol gefüllt ist 
	for _, character := range board[rowNumber]{
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
		if RowFull(board,i,symbol){
			return true 
		}
	}

	return false 

}

