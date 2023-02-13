package dreieckszahlen

// Dreieck berechnet für ein gegebenes n die n-te Dreieckszahl.
// Dies ist die Summe der Zahlen von 1 bis n.
func Dreieck(n int) int {
	// TODO

	if n <= 1 {
		return 1
	} 

	return n + Dreieck(n-1)
}

// Anmerkung:
// Diese Zahlen werden Dreieckszahlen genannt, weil jede dieser Zahlen eine
// mögliche Anzahl an Elementen (Kugeln, Bausteine o.Ä.) sind, aus denen man
// in der Fläche ein gleichseitiges Dreieck legen kann.
//
// Zum Beispiel kann man aus drei Kugeln ein Dreieck legen, für das nächstgrößere
// Dreieck braucht man dann sechs Kugeln, dann zehn usw.
