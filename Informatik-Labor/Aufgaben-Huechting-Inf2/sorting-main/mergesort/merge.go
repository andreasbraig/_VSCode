package mergesort

// Merge merges the two sorted sublists into a single sorted list.
/*func Merge(list1, list2 []int) []int {
	result := []int{}
	if len(list1) == 1 && len(list2) == 1 {

		el1 := list1[0]
		el2 := list2[0]

		if el1 <= el2 {
			return []int{el1, el2}
		}
		return []int{el2, el1}

	}
	if len(list1) == 1 {
		el1 := list1[0]
		el2 := list2[0]

		head2 := list2[len(list2)/2:]
		tail2 := list2[:len(list2)/2]

		if el1 <= el2 {
			result = append(result, Merge(head2, tail2)...)
			return result
		}

	}
	if len(list2) == 1 {
		el1 := list1[0]
		el2 := list2[0]

		head1 := list1[len(list1)/2:]
		tail1 := list1[:len(list1)/2]

		if el1 <= el2 {
			result = append(result, Merge(head1, tail1)...)
			return result
		}

	}

	head1 := list1[len(list1)/2:]
	tail1 := list1[:len(list1)/2]
	head2 := list2[len(list2)/2:]
	tail2 := list2[:len(list2)/2]

	result = append(result, Merge(head1, tail1)...)
	result = append(result, Merge(head2, tail2)...)

	return result
	// TODO+
}
*/
