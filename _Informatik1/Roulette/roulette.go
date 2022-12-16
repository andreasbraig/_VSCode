package main

import (
	"fmt"
	"math/rand"
)

func ReadNumber(msg string) int {

	//spieler wählt eine Zahl
	fmt.Print(msg)

	var input int
	fmt.Scanln(&input)

	return input

}

func main() {

	var input int = ReadNumber("Geben sie eine Zahl zwischen 0 und 36 ein: ")

	//Am Rad drehen

	var result int = rand.Int(36 - 0)

	// Auswerten ob spieler gewonnen hat

	if input == result {
		fmt.Println("Sie haben gewonnen!")
	} else {
		fmt.Println("Sie haben leider verloren.")
	}
}
