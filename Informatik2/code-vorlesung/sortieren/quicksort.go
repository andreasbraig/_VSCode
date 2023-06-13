package sortieren

func Quicksort(list []int) {

	if len(list) < 2 {
		return
	}

	// Wahl eines Pivot elements
	pivot := list[0]

	//Liste zerlegen
	left := []int{}
	right := []int{}

	for _, el := range list[1:] {
		if el < pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	//left und right auf magische Weise sortieren

	Quicksort(left)
	Quicksort(right)

	// left und right mit pivot wieder zusammensetzen
	result := []int{}
	result = append(left, pivot)
	result = append(result, right...)

	copy(list, result)
}
