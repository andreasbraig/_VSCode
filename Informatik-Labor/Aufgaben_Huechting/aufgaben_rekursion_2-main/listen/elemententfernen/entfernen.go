package elemententfernen

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
func RemoveElement(list []int, pos int) []int {
	// TODO

	if len(list) == 0 || pos < 0 {
		return list 
	}

	head, tail := list[0], list[1:]

	if pos == 0 {
		return tail 
	}
	// Hinweis:
	// Unterscheiden Sie zwei Fälle:
	// 1. Die Liste ist leer.
	// 2. Die Liste besteht aus einem Element und einer Folgeliste.
	//    Im zweiten Fall müssen Sie entscheiden, ob das erste Element gelöscht werden
	//    soll oder ein späteres.
	// Denken Sie zusätzlich daran, unsinnige pos-Angaben abzufangen.

	// nil ist ein Platzhalter für eine nicht-existente Liste.
	// Ersetzen Sie nil durch etwas sinnvolles.

	return append([]int{head},RemoveElement(tail, pos-1)...)


}
