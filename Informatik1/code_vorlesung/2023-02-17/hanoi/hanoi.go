package hanoi

import (
	"fmt"
	"strings"
)

// Problemlösung des Rätsels der Türme von Hanoi
//Hier soll ein Turm der größe h verschoben werden allerdings darf
//nie ein größerer Block auf einen kleineren Block gelegt werden.
//

func BewegePlatte(n int, start, Ziel string) {

	spaces := strings.Repeat("_", n)
	fmt.Printf("%s%s ==> %s\n", spaces, start, Ziel)

}

// Bewegt einen Turm der höhe 1
func Hanoi1(start, mitte, ziel string) {
	BewegePlatte(0, start, ziel)
}

// Bewegt einen Turm der Höhe 2
func Hanoi2(start, mitte, ziel string) {
	Hanoi1(start, ziel, mitte)
	BewegePlatte(0, start, ziel)
	Hanoi1(mitte, start, ziel)
}

// Bewegt einen turm der Höhe 3
func Hanoi3(start, mitte, ziel string) {
	Hanoi2(start, ziel, mitte)
	BewegePlatte(0, start, ziel)
	Hanoi2(mitte, start, ziel)
}

// Rekursive Problemlösung
// Man muss nur die Höhe der türme h übergeben
func Hanoi(h int, start, mitte, ziel string) {
	if h == 0 {
		return
	}
	Hanoi(h-1, start, ziel, mitte)
	BewegePlatte(h, start, ziel)
	Hanoi(h-1, mitte, start, ziel)
}
