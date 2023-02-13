package suche

// BinSearch geht davon aus, dass die Liste aufsteigend sortiert ist und führt
// eine sog. "binäre Suche" durch.
// D.h. die Funktion vergleicht zuerst das mittlere Element mit key.
// Falls key dort gefunden wurde, wird die Position zurückgeliefert.
// Ansonsten wird nur im linken oder rechten Teil weitergesucht.
// Ist das Element nicht enthalten, wird die Länge der Liste geliefert.
func BinSearch(list []int, key int) int {
	if len(list) == 0 {
		return 0
	}

	// TODO

	// Hinweis:
	// Bestimmen Sie zuerst die mittlere Position und vergleichen Sie das dortige
	// Element mit key. Danach fahren Sie je nach Wert des Elements rekursiv mit dem
	// linken und dem rechten Teil der Liste fort.

	mid := len(list) / 2

	left, right := list[:mid], list[mid+1:]

	if list[mid] == key {
		return mid
	}

	if list[mid] < key {
		return mid + 1 + BinSearch(right, key)
	}

	result := BinSearch(left, key)

	if result == len(left) {
		return len(list)
	}

	return result

}
