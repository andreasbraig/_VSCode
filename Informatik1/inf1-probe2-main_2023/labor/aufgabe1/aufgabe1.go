package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion LongestAbc.
PUNKTE: 10
BEWERTUNG:
*/

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {

// TODO

	filter := []string{}

// Eine Liste mit allen die mit abc beginnen erstellen

	for _, e := range list {

		if len(e) >= 3 {

			if string(e[0:3]) == "abc" {
			filter = append(filter, e)
			}

	}

	}

// längen vergleichen und elemente löschen.

//Den Fall der leeren Liste abfangen
	if len(filter) == 0 {
		return ""
	}

// den längsten string ermitteln
	max := 0
	result := ""

	for _, e := range filter {

		if len(e) > max {
			max = len(e)
			result = string(e)
			}

}

	return result

}
