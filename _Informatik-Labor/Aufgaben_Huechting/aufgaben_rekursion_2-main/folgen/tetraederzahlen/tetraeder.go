package tetraederzahlen

import "github.com/tel22a-inf/aufgaben_rekursion_2/rechnen/binomialkoeffizienten"

// Tetraeder berechnet für ein gegebenes n die n-te Tetraederzahl.
// Für die konkreten Zahlenwerte: siehe Tests.
func Tetraeder(n int) int {
	// TODO

	if n == 1 {
		return 1
	}

	return binomialkoeffizienten.BinomialCoefficient(n+2, 3)
}

// Anmerkung:
// Diese Zahlen sind die Fortsetzung der Dreieckszahlen für die dritte Dimenstion.
// Statt ein Dreieck zu legen, baut man einen Tetraeder aus Kugeln.
// Die Tetraederzahlen geben an, wie viele Kugeln man jeweils braucht.
