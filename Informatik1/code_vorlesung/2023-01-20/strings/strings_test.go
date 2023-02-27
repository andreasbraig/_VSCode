package strings

import "fmt"

// Anmerkung: Der Test schlägt noch fehl, obwohl er nicht sollte.
// Vermutlich liegt das an Leerzeichen, die nicht korrekt abgeglichen werden.
// Bugfix folgt.
func ExampleStrings() {

	// Erzeuge einen leeren String
	var s1 string
	fmt.Println(s1)

	// Erzeuge einen nicht-leeren String
	s2 := "Hallo"
	fmt.Println(s2)

	// Hänge etwas an einen String an:
	s2 += " "
	s2 = s2 + "Welt"
	fmt.Println(s2)

	// Den Buchstaben an Stelle 4 in s2
	// in eine Variable speichern:
	b1 := s2[4]

	// b1 ausgeben:
	fmt.Println(b1) // Gibt eine Zahl aus, den Unicode-Codepoint.
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n", s2)
	fmt.Println(string(b1))

	// b1 wieder als String auffassen und speichern.
	s3 := string(b1)
	fmt.Printf("%T\n", s3)
	fmt.Printf("%T\n", b1)

	// Schleife, die über s2 läuft und alle Buchstaben ausgibt.
	for pos, character := range s2 {
		fmt.Printf("%d: %c\n", pos, character)
		fmt.Println(pos, string(character))
	}

	// Output:
	//
	// Hallo
	// Hallo Welt
	// 111
	// uint8
	// string
	// o
	// string
	// uint8
	// 0: H
	// 0 H
	// 1: a
	// 1 a
	// 2: l
	// 2 l
	// 3: l
	// 3 l
	// 4: o
	// 4 o
	// 5:
	// 5
	// 6: W
	// 6 W
	// 7: e
	// 7 e
	// 8: l
	// 8 l
	// 9: t
	// 9 t
}
