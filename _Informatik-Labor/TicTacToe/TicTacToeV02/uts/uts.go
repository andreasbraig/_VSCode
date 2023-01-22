package uts

import (
	"fmt"
	"strings"
)

// Soll den spieler fragen, ob er noch eine runde spielen möchte.
//Gibt das in form eines Bools zurück

func AskforNewRound() bool {

	var input string

	fmt.Println("möchten Sie erneut spielen? Y/N")
	fmt.Scanln(&input)

	if strings.ToUpper(input) == "Y" { //abfrage nach neuer runde unabhängig der Groß und kleinschreibung
		fmt.Println("Die neue Runde startet...")
		return true
	}

	fmt.Println("Danke, dass sie TicTacToe by Andreas Braig gespielt haben!")
	return false

}
