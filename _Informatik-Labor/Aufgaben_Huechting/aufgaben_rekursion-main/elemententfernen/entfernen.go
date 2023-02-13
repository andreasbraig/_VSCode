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

	return append([]int{head}, RemoveElement(tail, pos-1)...)
}


