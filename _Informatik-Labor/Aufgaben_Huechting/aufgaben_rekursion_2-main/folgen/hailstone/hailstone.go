package hailstone

// Hailstone erwartet eine Zahl n und liefert eine Liste mit den
// Elementen der Hailstone-Folge ab n.
func Hailstone(n int) []int {
	// TODO

	if n == 1{
		return []int{1}
	}

	// Anmerkung: Das `result` oben ist mit n vorbelegt.
	// Die Idee ist, dass man nun entweder beendet, oder den nächsten Wert bestimmt,
	// die Hailstone-Folge dafür berechnet und an `result` anhängt.

	if n % 2 == 0 {
		return append([]int{n},Hailstone(n/2)...)
	}

	return append([]int{n},Hailstone(3*n+1)...)
}
