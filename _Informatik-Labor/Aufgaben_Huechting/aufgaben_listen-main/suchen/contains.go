package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Liefert true, falls die Liste die Zahl x enthält.
func Contains(list []int, x int) bool {
	// TODO

	for _, element := range list {

		if element == x {
			return true
		}

	}


	return false
}
