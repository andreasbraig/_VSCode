package main

import (
	"fmt"
)

//func Primedivide(n int) []int {

	//var primelist1 []int 

	//primelist1 = Primelist(n)


//}

func Primelist(n int) []int { //Funktion die eine Liste der Primzahlen bis n erstellt.

	var primelist []int // Variable in welcher die Primzahlen gespeichert werden

	for i := 2; i <= n; i++ { // Schleife die hochzählt
		if is_prime(i) == true { // Wenn eine der Zahlen der Schleife eine Primzahl ist wird sie in der Liste gespeichert
			primelist = append(primelist, i)
		}

	}

	return primelist

}

func main() {

	var input int

	input = import_int("Bitte geben sie einen Zahl ein:")

	fmt.Println(Primelist(input))

}
