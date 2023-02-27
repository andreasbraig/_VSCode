package roulette

import (
	"fmt"
	"math/rand"
	"time"
)

func ReadNumber(msg string) int {

	var input int

	//Programm überprüft, ob die Eingabe den Parametern entspricht
	for i := 0; i < 10; i++ { //Falls die Eingegebene Zahl nicht Stimmt, wird die Abfrage erneut durchgeführt

		fmt.Print(msg)
		//spieler wählt eine Zahl
		fmt.Scanln(&input)

		if input >= 0 && input <= 36 { //Abfrage
			return input
		}
		fmt.Println("Die Zahl war Ungültig!")

	}
	fmt.Println("Idiot!")
	return 37

}

// prüft ob der Spieler mit seiner gewählten Zahl gewonnen hat
// Braucht dazu die gewählte Zahl und die Zahl, die gedreht wurde
func PlayerWins(playerNumber, wheelNumber int) bool {

	if playerNumber == wheelNumber { //Bestimmt ob Player Number und Wheel Number dieselbe sind
		return true
	} else {
		return false
	}

}

// Gibt das ergebnis auf die Konsole aus
// Braucht dazu die gewählte Zahl und die Zahl, die gedreht wurde
func PrintResult(playerNumber, wheelNumber int) {

	fmt.Println("Es wurde", wheelNumber, "Gedreht")

	if PlayerWins(playerNumber, wheelNumber) { //Prüft ob der wert PlayerWins True oder false ist
		fmt.Println("Sie haben Gewonnen!")
	} else {
		fmt.Println("Sie haben leider verloren.")
	}
}

// dreht das roulette Rad
func SpinWheel() int {

	rand.Seed(time.Now().UnixNano())

	var number int = rand.Intn(37)

	return number

}

func main() {

	//Spieler Abfragen

	var input int = ReadNumber("Bitte geben sie eine Zahl zwischen 0 und 36 ein: ")

	//Am Rad drehen
	var result int = SpinWheel()

	// Auswerten ob spieler gewonnen hat

	PrintResult(input, result)

}
