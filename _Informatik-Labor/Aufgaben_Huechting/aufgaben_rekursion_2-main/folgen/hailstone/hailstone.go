package hailstone

// Hailstone erwartet eine Zahl n und liefert eine Liste mit den
// Elementen der Hailstone-Folge ab n.
func Hailstone(n int) []int {
	result := []int{n}
	// TODO

	if n == 1 {
		return result
	}


	if n % 2 == 0 {
		return append(result,Hailstone(n/2)...)
	}


	// Anmerkung: Das `result` oben ist mit n vorbelegt.
	// Die Idee ist, dass man nun entweder beendet, oder den nächsten Wert bestimmt,
	// die Hailstone-Folge dafür berechnet und an `result` anhängt.
	return append(result,Hailstone(3*n+1)...)
}
