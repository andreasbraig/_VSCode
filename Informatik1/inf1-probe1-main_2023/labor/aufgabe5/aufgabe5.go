package aufgabe5

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ArraySums.
PUNKTE: 15
BEWERTUNG:
*/

// ArraySums erwartet eine int-Slice l erwartet und liefert eine int-Slice,
// die an Stelle n die Summe der Elemente aus l bis zu Stelle n enthält.
func ArraySums(l []int) []int {
	result := []int{}
	// TODO

	// result an stelle 0 ist List an stelle 0 
		result = append(result, 0 )
	// result an stelle 1 ist List an stelle 1 plus Result an stelle 0

	for i,e := range l {
		result = append(result, result[i]+e)
	}

	// result an stelle 2 ist List an stelle 2 plus Resutl an stelle 1


	return result[1:]
}
