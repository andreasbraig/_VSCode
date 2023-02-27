package listenvergleich

// Liefert true, falls die beiden Listen gleich sind.
func ListsEqual(list1, list2 []int) bool {
	// TODO

	// Hinweis:
	// Unterscheiden Sie die folgenden Fälle:
	// 1. Die Listen sind beide leer.
	// 2. Eine der Listen ist leer, die andere nicht.
	// 3. Beide sind nicht leer.
	//

	if len(list1) != len(list2){
		return false
	}

	if len(list1) == 0 {
		return true
	}

	if list1[0] == list2[0] {
		return ListsEqual(list1[1:],list2[1:])
	}

	// In den beiden ersten Fällen können Sie sofort ein Ergebnis liefern,
	// im dritten Fall müssen Sie die beiden ersten Elemente vergleichen und dann
	// rekursiv mit den Rest-Listen fortfahren.

	return false
}
