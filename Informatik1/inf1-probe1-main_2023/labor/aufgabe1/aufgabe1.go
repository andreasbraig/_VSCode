package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MultFive.
PUNKTE: 5
BEWERTUNG:
*/

// MultFive erwartet eine Liste von Zahlen und liefert
// das Produkt der Elemente an den durch 5 teilbaren Positionen.
func MultFive(list []int) int {
	result := 1
	// TODO

	for i,e := range list {
		if i%5 == 0 {
			result *= e
		}
	}

	return result
}
