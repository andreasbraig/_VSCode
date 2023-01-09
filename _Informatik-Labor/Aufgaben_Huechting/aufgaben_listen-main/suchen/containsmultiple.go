package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Liefert true, falls die Liste x mehr als einmal enthält.
func ContainsMultiple(list []int, x int) bool {
	// TODO

	counter := 0

	for _, element := range list{

		if element == x {
			counter ++
		}

	}

	return counter > 1
}
