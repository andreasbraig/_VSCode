package sortieren

// BubbleSort sorts a list of integers
// It uses the bubble sort algorithm
func BubbleSort(list []int) {

	for i := 0; i < len(list); i++ {
		for j := 0; j+1 < len(list); j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

}

// BubbleUp is a helper function for BubbleSort2
// It bubbles up the largest element to the end of the list
func BubbleUp(list []int) bool {

	swapped := false

	// Die schleife läuft trotz des swapped wertes weiter bis das element an der richtigen stelle ist
	for j := 0; j+1 < len(list); j++ {
		if list[j] > list[j+1] {
			list[j], list[j+1] = list[j+1], list[j]
			swapped = true
		}
	}
	return swapped
}

// BubbleSort2 is a variant of BubbleSort with a function call to BubbleUp
func BubbleSort2(list []int) {
	for i := 0; i < len(list); i++ {
		if !BubbleUp(list) {
			return
		}
	}
}

// Eine weitere Alternative zu Bubble Sort -> Hier wird nur noch eine schleufe ausgeführt, bis bubble up false gibt. 
// Noch kürzer geschrieben, gleicher aufwand
func BubbleSort3(list []int) {
	for BubbleUp(list) {
	}
}
