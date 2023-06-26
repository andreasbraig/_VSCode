package insertionsort

import "github.com/tel22a-inf/sorting/bubblesort"

// MoveLeft moves the element at the given position to the left until it is in the correct position.
func MoveLeft(list []int, pos int) {
	// TODO
	//Begindend ab dem punkt des beginnens
	//Itteriere über die liste kleiner der position
	for i := pos; i > 0; i-- {
		//Überprüfe ob richtig sortiert
		if list[i] < list[i-1] {
			bubblesort.Swap(list, i, i-1) //Tausche nach hinten
		}
	}

}
