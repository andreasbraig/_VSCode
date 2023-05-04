package binomialkoeffizienten

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	// Hinweis: Berechnen Sie den Binomialkoeffizienten rekursiv.
	// Es gilt: (n,k) == (n-1,k-1) + (n-1,k)
	// Diese Berechnungsvorschrift können Sie direkt in eine Rekursion umsetzen,
	// Sie müssen nur noch den Rekursionsanfang hinzufügen.
	//
	// Anmerkung: Diese Berechnungsvorschrift beschreibt exakt das Pascalsche Dreieck!

	// TODO

	if k == 0 || k == n{
		return 1
	}

	// Hinweis:
	// Unterscheiden Sie, ob Sie am Rand oder in der Mitte des pascalschen Dreiecks
	// sind. Bei n == 0, k==0 oder k==n ist das Ergebnis sofort klar, ansonsten
	// können Sie rekursiv weitermachen.

	return BinomialCoefficient(n-1,k-1) + BinomialCoefficient(n-1,k)
}
