package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion GetStringsBetween.
PUNKTE: 15
BEWERTUNG:
*/

// GetStringsBetween eine Liste und zwei Strings first und last.
// Die Funktion liefert die Teilliste, die zwischen first und last liegt.
// first sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func GetStringsBetween(list []string, first, last string) []string {
	result := []string{}
	// TODO


	// Iteriert über die Liste Range 
	// -> Die schleife bricht ab wenn die Liste fertig ist

	for i,e := range list {

		if string(e) == first {
			result = append(result, list[i+1:]...)
		}

	}

	// -> Alles was nach Begin ist wird in die Liste eingefügt 


	//Iteriert über Result und entfernt alles was nach dem wort End steht inklusive dem Wort selbst 





	return result
}
