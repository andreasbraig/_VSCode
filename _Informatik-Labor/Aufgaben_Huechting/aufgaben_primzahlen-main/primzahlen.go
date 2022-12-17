package main

import "fmt"

func main() {}

// Erwartet zwei Zahlen m und n.
// Liefert true, falls m ein Teiler von n ist.
func Divides(m, n int) bool {
	// Hinweis: Es ist eine Lösung vorgegeben, die man auch in der Praxis verwenden würde.
	// D.h. eigentlich würde man diese Funktion nicht schreiben
	// oder sie so einfach lassen, wie hier in der Vorgabe.
	//
	// Als Übungsaufgabe ersetzen Sie diese Lösung dennoch durch eine,
	// die den Modulo-Operator nicht verwendet.
	if m == 0 {
		return false
	}

	var result int = n / m

	if result*m == n {
		return true
	}
	
	return false
	// Erläuterung:
	// Der Modulo-Operator % liefert den Rest der ganzzahligen Division n/m.
	// Wenn dieser Rest 0 ist, dann ist n durch m teilbar.
}

// Erwartet eine Zahl n.
// Liefert true, falls n eine Primzahl ist.
func IsPrime(n int) bool {
	// TODO
	if n <= 1 { //Alle Zahlen kleiner gleich 1 sind keine Primzahlen und müssen nicht weiter überprüft werden
		return false
	} else {
		for i := 2; i < n; i++ { //for schleife, die mögliche Weitere Teiler größer 1 Probiert
			if n%i == 0 { //Sobald ein Teiler gefunden wurde, ist es keine Primzahl
				return false
			}
		}

	}

	return true
}

// Erwartet eine Zahl n.
// Gibt alle Primzahlen auf der Konsole aus, die kleiner als n sind.
func PrintPrimes(n int) {
	// TODO
	for i := 2; i < n; i++ {
		if IsPrime(i) {
			fmt.Println(i)
		}
	}

}

// Erwartet eine Zahl n.
// Liefert die nächstgrößere Primzahl.
// Liefert n, falls n selbst eine Primzahl ist.
func NextPrime(n int) int {
	// TODO
	for IsPrime(n) == false {
		if IsPrime(n) {
			return n
		} else {
			n++
		}
	}
	return n

}

// Erwartet eine Zahl n.
// Liefert den nächsten Primzahlzwilling, der größer ist als n.
// Genauer: Liefert die kleinste Zahl k, so dass:
// * k >= n
// * k ist eine Primzahl
// * k + 2 ist eine Primzahl
func NextPrimeTwin(n int) int {
	// TODO
	return 0
}

// Erwartet eine Zahl n.
// Liefert die größte Primzahl, die echt kleiner als n ist.
// Falls es keine solche Zahl gibt, wird 0 geliefert.
func GreatestPrimeBelow(n int) int {
	// TODO
	for IsPrime(n) == false {
		n--
		if IsPrime(n) {
			return n
		} else {
		}
	}
	return 0
}
