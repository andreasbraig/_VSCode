package distance

import "fmt"

// ExampleDistance_constructors erzeugt verschiedene Distances jeweils mit
// einem der verschiedenen Konstruktoren.
// Jedes erzeugte Objekt wird unmittelbar danach ausgegeben.
func ExampleDistance_constructors() {
	// Erzeuge jeweils ein Distance-Objekt mit einem der Konstruktoren und gebe es aus:

	d1 := FromMM(1)
	fmt.Println(d1)

	d2 := FromCM(1)
	fmt.Println(d2)

	d3 := FromM(1)
	fmt.Println(d3)

	d4 := FromKM(1)
	fmt.Println(d4)

	d5 := FromMI(1)
	fmt.Println(d5)

	// In den Ausgaben unten sehen wir, dass das Struct
	// immer eine Millimeter-Anzahl enthält.

	// Output:
	// 1
	// 10
	// 1000
	// 1000000
	// 1609340
}
