package main

import (
	"fmt"
)

func is_prime(n int) bool { // Funktion für die Überprüfung der Primzahlen

	if n <= 1 {
		return false //alle Zahlen kleiner gleich 1 müssen nicht auf Teilbarkeit überprüft werden
	} else {
		for i := 2; i < n; i++ { // probiere alle Zahlen die kleiner sind als die angegebene Zahl n auf Teilbarkeit mit der Zahl n
			if n%i == 0 {
				return false //Sobald mehr als ein teiler gefunden wurde ist die Zahl keine Primzahl
			}
		}
		return true
	}

}

func import_int(msg string) int { //Funktion der man einen satz votgibt, die einen Int abfragt 

	var input int

	fmt.Print(msg)
	fmt.Scanln(&input)  //Abfrage für Input 
	fmt.Println("Sie haben", input, "eingegeben")

	return input
}
