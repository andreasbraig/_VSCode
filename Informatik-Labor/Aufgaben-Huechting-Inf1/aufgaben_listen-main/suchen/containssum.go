package suchen

// Erwartet eine Liste von Zahlen.
// Liefert true, falls die Liste an irgendeiner
// Stelle eine Kette von drei Zahlen enthält,
// deren Summe 42 ist.
func ContainsSum(list []int) bool {
	// TODO

	if len(list) < 3 {
		return false
	}

	for idx, element := range list[:len(list)-2] { //die -2 gibt an wie viel man "vorschauen" will
		if element + list[idx+1] + list [idx +2] == 42 {
			return true
		}
	}

	return false
}
