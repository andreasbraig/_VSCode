package selectionsort

// SmallestPos returns the index of the smallest element in the given list.
func SmallestPos(list []int) int {
	smallest := 0
	smel := list[0]
	// TODO
	for i, el := range list {
		if el < smel {
			smallest = i
			smel = el
		}
	}
	return smallest
}
