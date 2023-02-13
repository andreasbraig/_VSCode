package multiplikation

// Liefert das Produkt von x und y.
func Mult(x, y int) int {
	// TODO

	// Hinweis:
	// Unterscheiden Sie zwei Fälle:
	// 1. y == 0
	// 2. y != 0
	//
	// Im zweiten Fall müssen Sie die Multiplikation mittels Rekursion
	// auf die Addition zurückführen.

	if y == 0{
		return 0
	}

	return x + Mult(x,y-1)
}
