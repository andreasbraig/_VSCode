package partition

import "github.com/tel22a-inf/aufgaben_rekursion/filter"

// Liefert zwei Listen:
// - Eine, die alle Elemente aus list enthält, die kleiner oder gleich key sind.
// - Eine, die alle übrigen Elemente aus list enthält.
func Partition(list []int, key int) ([]int, []int) {
	// TODO
	return filter.FilterLess(list,key), filter.FilterGreater(list,key)
}

/* Hinweis:
Diese Funktion direkt rekursiv zu schreiben, ist etwas komplizierter.
Lösen Sie zuerst die Filter-Aufgaben und verwenden Sie dann die dortigen Funktionen.
Die Funktion Partition selbst muss dann nicht rekursiv sein.
*/