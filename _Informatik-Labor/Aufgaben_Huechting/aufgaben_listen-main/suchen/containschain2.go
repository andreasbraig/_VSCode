package suchen

// Erwartet eine Liste von Zahlen.
// Liefert true, falls die Liste an irgendeiner
// Stelle eine aufsteigende Kette von mindestens
// drei Zahlen enthält.
// Im Gegensatz zu `ContainsChain()` müssen die Zahlen
// hier nicht direkt aufeinanderfolgen.
func ContainsChain2(list []int) bool {
	// TODO

	if len(list) < 3 {
		return false
	}

	counter := 1

	for idx, element := range list[:len(list)-1] { // damit an der eins höheren stelle geschaut werden kann muss die position um eins veringert werden

		if element < list[idx+1] { // element muss gleich der Liste an der stelle stelle[idx+1]  minus eins sein

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
