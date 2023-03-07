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

	// TODO

	firstpos := -1
	lastpos := -1

	for i,e :=range list {
		if e == first {
		
			firstpos = i

		}
		if e == last {
			lastpos = i
		}

	}

	if firstpos == -1 || lastpos <= firstpos {
		return []string{}
	} 

	return list[firstpos+1:lastpos]
}
