package quicksort

// QuickSort sorts the given list using the quicksort algorithm.
func QuickSort(list []int) {
	// TODO

	//Rekursionanker
	if len(list) < 2 {
		return
	}

	//select Pivot

	pivot := list[0]
	//leere Listen erstellen
	left := []int{}
	right := []int{}

	//Listen aufteilen
	for _, el := range list[1:] {
		if el < pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	//Rekursionsaufruf um teillisten zu sortieren

	QuickSort(right)
	QuickSort(left)

	//alles zusammenführen

	result := append(left, pivot)
	result = append(result, right...)

	//copy überschreibt die liste mit dem result
	copy(list, result)

}
