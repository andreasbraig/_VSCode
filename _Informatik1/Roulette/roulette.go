package main

import (
	"fmt"
)

func ReadNumber(msg string) int {

	//spieler wählt eine Zahl
	fmt.Print(msg)

	var input int
	fmt.Scanln(&input)

	return input

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

	if PlayerWins(playerNumber, wheelNumber) { //Prüft ob der wert PlayerWins True oder false ist
		fmt.Println("Sie haben Gewonnen!")
	} else {
		fmt.Println("Sie haben leider verloren.")
	}
}

func main() {

	//Spieler Abfragen

	var input int = ReadNumber("Bitte geben sie eine Zahl zwischen 0 und 36 ein: ")

	//Am Rad drehen

	var result int = 13

	// Auswerten ob spieler gewonnen hat

	PrintResult(input, result)

}
