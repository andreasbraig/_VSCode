package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Falls es eine Stelle gibt, an der x zwei Mal
// hintereinander vorkommt, liefert die Funktion
// diese Stelle.
// Liefert -1, falls die Situation nicht auftritt.
func FindPair(list []int, x int) int {
	// TODO

	if len(list) < 3 {
		return -1
	}

	for idx, element := range list[:len(list)-1] { // damit an der eins höheren stelle geschaut werden kann muss die position um eins veringert werden

		if element == x && list[idx+1] == x { // element muss gleich der Liste an der stelle stelle[idx+1]  minus eins sein

		return idx

		}

	}

	return -1
}
