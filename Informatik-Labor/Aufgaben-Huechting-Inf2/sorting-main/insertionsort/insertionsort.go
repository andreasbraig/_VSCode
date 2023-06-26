package insertionsort

// InsertionSort sorts the given list using the insertion sort algorithm.
func InsertionSort(list []int) {
	// TODO

	for i := range list {
		MoveLeft(list, i)
	}

}
