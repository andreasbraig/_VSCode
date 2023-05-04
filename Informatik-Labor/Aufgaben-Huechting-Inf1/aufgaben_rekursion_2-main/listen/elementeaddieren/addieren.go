package elementeaddieren

// Liefert die Summe aller Elemente in list.
func AddElements(list []int) int {
	// TODO

	if len(list) == 0 {
		return 0
	}
	
	head, tail := list[0], list[1:]

	return head + AddElements(tail)

	// Hinweis:
	// Unterscheiden Sie zwei Fälle:
	// 1. Die Liste ist leer.
	// 2. Die Liste besteht aus einem Element und einer Folgeliste.
}
