package filter

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
func FilterLess(list []int, key int) []int {
	// TODO
	if len(list) == 0 {
		return list
	}

	head, tail := list[0], list[1:]

	if head > key {

		return FilterLess(tail, key)

	}

	return append([]int{head}, FilterLess(tail, key)...)

}

// Liefert eine Liste mit allen Elementen aus list, die größer als key sind.
func FilterGreater(list []int, key int) []int {
	// TODO
	if len(list) == 0 {
		return list
	}

	head, tail := list[0], list[1:]

	if head < key {

		return FilterGreater(tail, key)

	}

	return append([]int{head}, FilterGreater(tail, key)...)
}
