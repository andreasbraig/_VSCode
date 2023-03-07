package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Power2.
Die Funktion muss rekursiv sein.
PUNKTE: 5
BEWERTUNG:
*/

// Power2 erwartet einen int-Parameter x und liefert die Potenz "2 hoch x".
func Power2(x int) int {
	// TODO
	if x == 0 {
		return 1
	}
	return 2 * Power2(x-1)
}
