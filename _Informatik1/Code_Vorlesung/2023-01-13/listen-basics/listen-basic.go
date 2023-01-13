package main

import "fmt"

func main() {

	//Eine Variable definieren, in der Zahlen gespeichert werden könen

	var zahl1 int

	// diese Zahl ausgeben

	fmt.Println(zahl1)

	// Eine Variable direkt mittels Wert definierne

	zahl2 := 42

	fmt.Println(zahl2)

	// Eine Variable für eine Liste von Zahlen definieren:

	var list1 []int

	fmt.Println(list1)

	// list 1 ist noch leer. Wir hängen ein Element an:

	list1 = append(list1, 55)

	fmt.Println(list1)

	list1 = append(list1, 42, 23, 38)

	fmt.Println(list1)

	//Eine zweite Liste erzeugen:
	list2 := []int{11, 22, 33, 44, 55}
	fmt.Println(list2)

	list1 = append(list1, list2...) // der ... bedeutet die liste auspacken und deren INT werte als einzelne zahlen anzuhängen
	fmt.Println(list1)

	// Das Element an stelle zwei von list1 Ausgeben
	fmt.Println(list1[2])

	// alle Elemente in list1 in einer Schleife ausgeben.
	for i := 0; i < len(list1); i = i + 1 {
		fmt.Println(i, ",", list1[i])
	}

	//Alle Elemente in list1 in einer Schleife über Range ausgeben

	for idx, element := range list1 { //erste Variable (idx) ist das "i" und zweite Variable (element) ist list1[i] liste an stelle i
		fmt.Println(idx, ":", element)
	}

	//Jedes Element in list 1 verdoppeln 

	for i := 0; i < len(list1); i++{
		list1[i] *= 2 // ist gleich wie list1[i] = list1[i] *2
	}
	fmt.Println(list1)
	
	// Über Range befehl gelöst: 
	
	for idx, element := range list1 { //erste Variable (idx) ist das "i" und zweite Variable (element) ist list1[i] liste an stelle i
		
		list1[idx] = element * 2
		
	}
	fmt.Println(list1)

}
