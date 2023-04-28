package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountGreater.
Die Funktion muss rekursiv sein.
PUNKTE: 10
BEWERTUNG:
*/

// CountGreater erwartet eine Liste von Zahlen und eine Zahl x.
// Sie liefert die Anzahl der Zahlen darin, die größer als x sind.
func CountGreater(list []int, x int) int {
	// TODO

	if len(list) == 0 {
		return 0
	}

	head, tail := list[0] , list[1:]

	if head > x {
		return 1 + CountGreater(tail,x)
	}

	return CountGreater(tail,x)
}
