package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountEven.
Die Funktion muss rekursiv sein.
PUNKTE: 10
BEWERTUNG:
*/

// CountEven erwartet eine Liste von Zahlen und liefert die Anzahl der geraden Zahlen darin.
func CountEven(list []int) int {
	// TODO

	//anker wenn die liste leer ist
	if len(list)== 0 {
		return 0
	}

	// aufbrechen in Head und Tail 

	// Fall1: Head ist gerade 
	// -> Head aufsummieren Tail weitergeben 

	if list[0]%2 == 0 {
		return 1 + CountEven(list[1:])
	}


	// Fall2: Head ist ungersde 
	// -> Head ignorieren und Tail weitergeben. 


	return CountEven(list[1:])
}
