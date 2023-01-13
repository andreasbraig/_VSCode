package main

import (
	"fmt"
)

func main() {

	// Spieler Wählt eine Zahl
	var input int

	fmt.Print("Geben sie eine Zahl zwischen 0 und 36 ein. ")
	fmt.Scanln(&input)
	fmt.Println("Rien ne va plus! Sie haben ", input, "eingegeben.")

	//Am Rad Drehen
	var result int = 13

	//Auswerten ob der Spieler gewonnen hat

	fmt.Println("es wurde ", result, "Gewürfelt")

	if input == result {
		fmt.Println("Sie haben Gewonnen!")
	} else {
		fmt.Println("Sie haben leider verloren.")
	}
}
