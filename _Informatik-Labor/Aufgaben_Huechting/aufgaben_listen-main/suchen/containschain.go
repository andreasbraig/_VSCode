package suchen

// Erwartet eine Liste von Zahlen.
// Liefert true, falls die Liste an irgendeiner
// Stelle eine aufsteigende Kette von mindestens
// drei aufeinanderfolgenden Zahlen enthält.
func ContainsChain(list []int) bool {
	// TODO

	//abfangen der zu kurzen Listen

	if len(list) < 3 {
		return false
	}

	counter := 1

	for idx, element := range list[:len(list)-1] { // damit an der eins höheren stelle geschaut werden kann muss die position um eins veringert werden

		if element == list[idx+1]-1 { // element muss gleich der Liste an der stelle stelle[idx+1]  minus eins sein

			counter++

		} else {
			counter = 1
		}

		if counter >= 3 {
			return true
		}

	}

	return false
}
