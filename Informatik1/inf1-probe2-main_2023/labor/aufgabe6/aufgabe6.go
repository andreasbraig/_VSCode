package aufgabe6

/*
AUFGABENSTELLUNG:
Hier ist ein Struct Person vorgegeben.
Vervollständigen Sie die Funktion CommonFriends.
PUNKTE: 15
BEWERTUNG:
*/

// Person repräsentiert eine Person in einem sozialen Netzwerk.
type Person struct {
	Name    string
	Friends []Person
}

// CommonFriends erwartet zwei Personen und liefert eine Liste
// mit allen Personen, die mit beiden befreundet sind.
func CommonFriends(p1, p2 Person) []Person {
	result := []Person{}
	// TODO

	for _,friendp1 := range p1.Friends {
		for _,friendp2 := range p2.Friends {
			if friendp1.Name == friendp2.Name {
				result = append(result, friendp1)
			}
		}
	}

	return result
}
