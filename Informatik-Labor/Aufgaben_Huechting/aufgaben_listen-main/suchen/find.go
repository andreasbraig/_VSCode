package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Liefert die Position von x in der Liste.
// Liefert -1, falls x nicht enthalten ist.
func Find(list []int, x int) int {
	// TODO

	for idx, element := range list {
		if element == x {
			return idx
		}
	}

	return -1
}
