package specialfeature

//Unused
//Soll ein neues, Breiteres spielfeld erstellen

func NewBoard(n int) [][]string {

	board := [][]string{}
	field := []string{}

	for m := 0; m < n; m++ {
		
		field = append(field, " ")
	
	}

	for i := 0; i < n; i++ {

		board = append(board, field)

	}
	return board

}
