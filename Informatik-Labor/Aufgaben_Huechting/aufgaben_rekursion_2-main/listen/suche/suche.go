package suche

// Search führt eine lineare Suche durch.
// D.h. es durchsucht die Liste von Anfang bis ende nach key und liefert die Position,
// an der Key steht. Liefert die Länge der Liste, falls key nicht enthalten ist.
func Search(list []int, key int) int {

	// TODO

	// Hinweis:
	// Unterscheiden Sie drei Fälle:
	// 1. Die Liste ist leer.
	if len(list) == 0 {
		return 0 
	}
	// 2. Das erste Element ist der gesuchte Wert.
	if list[0] == key {
		return 0
	}

	// 3. Das erste Element ist nicht der gesuchte Wert.
	return 1 + Search(list[1:],key)
	// In den beiden ersten Fällen sind Sie fertig, im dritten Fall müssen Sie
	// rekursiv im Rest der Liste weitersuchen.
}
