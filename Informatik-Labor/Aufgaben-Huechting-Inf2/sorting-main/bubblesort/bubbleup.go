package bubblesort

// BubbleUp implements a single iteration of the bubble sort algorithm.
// It iterates over the list and compares each element with the next one.
// If the element is larger than the next one, they are swapped.
// The function returns true if at least one swap was performed.
func BubbleUp(list []int) bool {
	swapped := false
	// TODO

	for i,e := range list[:len(list)-1] {
		e1 := e
		e2 := list[i+1]
		if e1 > e2 {
			Swap(list,i, i+1)
			swapped = true
		}
	}

	return swapped
}


func Swap(list []int,i1 ,i2 int) {

	list[i1] ,list[i2] = list[i2], list[i1]

	 
}