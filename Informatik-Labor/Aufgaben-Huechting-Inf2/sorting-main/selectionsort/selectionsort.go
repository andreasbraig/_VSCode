package selectionsort

// SelectionSort sorts the given list using the selection sort algorithm.
func SelectionSort(list []int) {
	// TODO

	for i := range list {

		Swap(list[i:], SmallestPos(list), 0)

	}

}

func Swap(list []int, i1, i2 int) {

	list[i1], list[i2] = list[i2], list[i1]

}
