package suchen

// Erwartet eine Liste von Zahlen.
// Liefert die Position der ersten negativen Zahl in der Liste.
// Liefert -1, falls keine negative Zahl enthalten ist.
func FindNegative(list []int) int {
	// TODO

	for idx, element := range list {
		if element < 0 {
			return idx
		}
	}

	return -1
}
