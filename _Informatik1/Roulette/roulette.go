package main

import (
	"fmt"
	"math/rand"
	"time"
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
