package aufgabe7

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Difference.
PUNKTE: 10
BEWERTUNG:
*/

// Difference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen aus l1,
// die nicht in l2 vorhanden sind.
func Difference(l1, l2 []int) []int {
	result := []int{}
	// BEGIN-SOLUTION

	// TODO

	//Iteriert über die Liste l1 und schaut ob das element in l2 vorkommt, wenn nicht wird result erweitert.

	for _, el1 := range l1 {

		filter := []int{}

		for _, el2 := range l2 {

			if el1 == el2 {
				filter = append(filter, el1)
			}

		}

		if CheckListLen(filter) {
			result = append(result, el1)
		}

	}

	// END-SOLUTION
	return result
}

// liefert true wenn die liste länger 0 ist
func CheckListLen(list []int) bool {

	return len(list) == 0

}
