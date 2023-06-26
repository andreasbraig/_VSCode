package mergesort

// Merge merges the two sorted sublists into a single sorted list.
func Merge(list1, list2 []int) []int {
	result := []int{}
	// TODO+

	swapped := false
	counter := 0

	min := len(list1)
	if len(list2) < len(list1) {
		min = len(list2)
		swapped = true 	
	}

	for i := 0; i< min ; i++{

		result = append(result, list1[i])
		result = append(result, list2[i])
		counter ++
	}

	if swapped {
		result = append(result, list1[counter:]... )
	}
	result = append(result, list2[counter:]... )

	return result
}


