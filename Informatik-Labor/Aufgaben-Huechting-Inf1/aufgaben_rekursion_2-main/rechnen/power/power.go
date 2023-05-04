package power

// Liefert die Potenz "x hoch y".
func Power(x, y int) int {
	// TODO

	// Hinweis:
	// Gehen Sie analog zur Multiplikation vor.

	if y == 0 {
		return 1
	}

	return x * Power(x,y-1)
}


