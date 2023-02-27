package velocity

import "fmt"

// ExampleVelocity_getters erzeugt ein Velocity-Objekt als km/h
// und konvertiert es nach m/s und wieder zurück.
func ExampleVelocity_getters() {
	v1 := FromKMH(75)
	v1_mps := v1.ToMPS()
	v1_roundtrip := FromMPS(v1_mps).ToKMH()
	fmt.Println(v1)
	fmt.Println(v1_mps)
	fmt.Println(v1_roundtrip)

	// Anmerkung: Wir erwarten unten als Ergebnis des RoundTrip 72 und nicht 75.
	// Das liegt an Rundungsfehlern. Wie kann man das beheben/verbessern?

	// Output:
	// {75000000 3600000}
	// 20
	// 72
}
