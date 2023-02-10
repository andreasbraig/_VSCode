package listenvergleich

// Liefert true, falls die beiden Listen gleich sind.
func ListsEqual(list1, list2 []int) bool {
	// TODO
	
	if len(list1) != len(list2) {
		return false
	}

	if len(list1) == 0 || len(list2) == 0{
		return true

	}

	head1,tail1 := list1[0],list1[1:]
	head2,tail2 := list2[0],list2[1:]

	if head1 == head2 {
		return ListsEqual(tail1,tail2)
	}
	
	return false
}
