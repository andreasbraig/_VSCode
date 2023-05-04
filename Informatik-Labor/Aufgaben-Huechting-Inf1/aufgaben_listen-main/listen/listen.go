package listen

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie oft n in der Liste vorkommt.
func CountElement(list []int, n int) int {
	// ANMERKUNG: Dies ist ein Beispiel, hier ist nichts zu tun.
	result := 0
	for _, element := range list {
		if element == n {
			result++
		}
	}
	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie viele Elemente in der Liste größer als n sind.
func CountGreaterThan(list []int, n int) int {
	result := 0
	for _, element := range list {
		// TODO: Diese Anweisung ist ein Platzhalter, damit der Code compiliert.
		//       Ersetzen Sie sie durch die richtige Lösung.
		if element > n {
			result ++
		}
	}
	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie viele Elemente in der Liste ungleich n sind.
func CountNotN(list []int, n int) int {
	result := 0
	for _, element := range list {
		// TODO: Diese Anweisung ist ein Platzhalter, damit der Code compiliert.
		//       Ersetzen Sie sie durch die richtige Lösung.
		if element != n {
			result ++
		}
	}
	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie viele Elemente in der Liste durch n teilbar sind.
func CountDivisors(list []int, n int) int {
	result := 0
	// TODO
	for _, element := range list {
		// TODO: Diese Anweisung ist ein Platzhalter, damit der Code compiliert.
		//       Ersetzen Sie sie durch die richtige Lösung.
		if element % n == 0 {
			result ++
		}
	}
	return result
}

// Erwartet zwei Listen.
// Liefert zurück, an wie vielen Stellen die Elemente in den beiden Listen gleich sind.
func CountEqualElements(list1, list2 []int) int {
	// TODO: Diese Lösung funktioniert für gleich lange Listen, aber nicht für
	// unterschiedlich lange Listen. Es gibt zwei Tests zu dieser Funktion,
	// einer ist ok, der andere noch nicht.
	result := 0

	minlist := list1 // Definiert eine kürzeste Liste

	if len(minlist) > len(list2) { // Überprüft ob diese Liste kürzer als die andere ist 
		
		minlist = list2

	}

	for i := range minlist {      //Für die Länge "Range" der Kürzeren Liste werden beide Listen verglichen
		if list1[i] == list2[i] {
			result++
		}
	}
	return result
}

// Erwartet zwei Listen.
// Liefert zurück, an wie vielen Stellen die Elemente in list1 größer
// als in list 2 sind. Sind die Listen nicht gleich lang, soll die Stelle ignoriert werden.
func CountGreaterElements(list1, list2 []int) int {
	result := 0
	// TODO: Auch hierfür gibt es zwei Tests, die beide ok sein müssen.


	minlist := list1 // Definiert eine kürzeste Liste

	if len(minlist) > len(list2) { // Überprüft ob diese Liste kürzer als die andere ist 
		
		minlist = list2

	}

	for i := range minlist {      //Für die Länge "Range" der Kürzeren Liste werden beide Listen verglichen
		if list1[i] > list2[i] {
			result++
		}
	}

	return result
}
