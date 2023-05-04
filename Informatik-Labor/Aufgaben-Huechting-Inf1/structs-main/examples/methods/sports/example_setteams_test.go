package sports

import "fmt"

// SetHomeTeam erwartet einen String und schreibt ihn in das Attribut Home.
// Diese Methode hat einen Pointer-Receiver.
// D.h. wenn sie aufgerufen wird, liegt m als Pointer auf ein Match vor.
// Dies bedeutet, dass sich Änderungen als *Seiteneffekt* auf das Objekt auswirken,
// auf dem die Methode aufgerufen wurde.
func (m *Match) SetHomeTeam(name string) {
	m.Home = name
}

// ExampleMatch_SetHomeTeam demonstiert, dass der Aufruf von SetHomeTeam auf einem
// Match-Objekt den gewünschten Effekt hat.
func ExampleMatch_SetHomeTeam() {
	// Erzeuge ein Match-Objekt:
	m := Match{Home: "Keiner", Visitors: "Niemand", Location: "Mannheim", Result: 0}

	// Schreibe einen Namen ins Home-Feld mittels der Methode SetHomeTeam:
	m.SetHomeTeam("Wir")

	// m ausgeben:
	fmt.Println(m)

	// Output:
	// {Wir Niemand Mannheim 0}
}

// SetVisitorTeam erwartet einen String und schreibt ihn in das Attribut Visitors.
// Diese Methode hat keinen Pointer-Receiver.
// D.h. wenn sie aufgerufen wird, liegt m nicht als Pointer auf ein Match vor,
// sondern als Kopie des Matches, auf dem die Funktion aufgerufen wurde.
// Dies bedeutet, dass sich Änderungen nicht auf das Objekt auswirken,
// auf dem die Methode aufgerufen wurde.
func (m Match) SetVisitorTeam(name string) {
	// Hier wird zwar m.Visitors geschrieben, dies hat aber keinen Effekt.
	m.Visitors = name

	// Genauer: Wenn man diese Funktion im Debugger laufen lässt,
	// sieht man, dass zwar das Attribut des Objekts m geändert wird.
	// Nach Ende der Funktion sieht man die Änderung außen aber nicht mehr.
}

// ExampleMatch_SetVisitorTeam demonstiert, dass der Aufruf von SetHomeTeam auf einem
// Match-Objekt den gewünschten Effekt hat.
func ExampleMatch_SetVisitorTeam() {
	// Erzeuge ein Match-Objekt:
	m := Match{Home: "Keiner", Visitors: "Niemand", Location: "Mannheim", Result: 0}

	// Schreibe einen Namen ins Home-Feld mittels der Methode SetHomeTeam:
	m.SetVisitorTeam("Die anderen")

	// m ausgeben:
	fmt.Println(m)

	// Man sieht, dass der Aufruf keinen Effekt hatte:
	// Output:
	// {Keiner Niemand Mannheim 0}
}
