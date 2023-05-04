package sortieren

import "github.com/tel22a-inf/aufgaben_rekursion/partition"

func QuickSort(list []int) []int {
	//TODO

	if len(list) == 0 { //Rekursionsanker
		return list
	}

	pivot, tail := list[0], list[1:] // Aufeilen in das erste Pivot element und dem Rest der Liste

	//Liefert Zwei Listen eine Größer als Pivot eine kleiner gleich privot
	res1, res2 := partition.Partition(tail, pivot) // Filtern der Restlichen liste mit Pivot als Sortierelement

	// Neusortierung der beiden Listen nach obigem Algorithmus
	res1 = QuickSort(res1)
	res2 = QuickSort(res2)

	//Zusammmensetzen der beiden Sortierten listen mit pivot als mittelpunkt
	result := append(res1, pivot)
	result = append(result, res2...)

	return result
}
