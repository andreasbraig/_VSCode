package main

func main() {}

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	// TODO
	for i := 0; i < len(s); i++ { // schleife läuft über die länge von s
		if s[i] == c { //sobald Buchstabe c in s an stelle i gefunden wurde return true
			return true
		}
	}
	return false
}

// Erwartet einen String s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c byte) int {
	count := 0
	for i := 0; i < len(s); i++ { // schleife läuft über die länge von s
		if s[i] == c { //sobald Buchstabe c in s an stelle i gefunden wurde Zähle den counter um eins hoch
			count++
		}
	}
	return count //übergebe den Counter zurück
}

// Erwartet einen String s und einen Buchstaben c.
// Liefert die Position, an der c in s vorkommt.
// Liefert die Länge von s, falls c nicht in s vorkommt.
// Kommt c mehrfach vor, soll die erste Position geliefert werden.
func PositionOf(s string, c byte) int {

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return len(s)
}

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {

	l1 := len(s)
	l2 := len(t)

	for i := 0; i <= l1-l2; i++ {
		if s[i:i+l2] == t {
			return true
		}

	}
	return false

}

// Erwartet einen String und prüft, ob er korrekte Klammerpaare enthält.
// D.h. der Eingabestring enthält Klammern '(' und ')'.
// Die Funktion soll nun prüfen, ob der String für jede öffnende Klammer auch eine
// schließende Klammer enthält.
// Außerdem darf es keine schließenden Klammern geben, für die es nicht vorher eine
// passende öffnende Klammer gegeben hat.
// Die Funktion soll true liefern, falls der String korrekt geklammert ist.
func CheckParentheses(s string) bool {
	counteropen := 0
	counterclose := 0
	// TODO
	for i := 0; i < len(s); i++ { // schleife läuft über die länge von s
		if s[i] == '(' { //sobald Klammer auf in s an stelle i gefunden wurde counter steigt
			counteropen++
		}
	}

	for i := 0; i < len(s); i++ { // schleife läuft über die länge von s
		if s[i] == ')' { //sobald Klammer zu in s an stelle i gefunden wurde counter steigt
			counterclose++
		}
	}
	return counteropen == counterclose //vergleiche anzahl der Offnenen mit geschlossenen Klammern, wenn gleich bedingung erfüllt
}

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {
	result := ""
	// TODO
	for i := 1; i < n; i++ {
		result = result + s + sep
	}

	return result + s
}

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip1(s1, s2 string) string {
	result := ""
	// TODO

	// Prüfung welcher der Strings der längere ist
	var max int
	if len(s1) > len(s2) {
		max = len(s1)
	} else {
		max = len(s2)
	}

	//Zusammenbau der Strings zu einem String
	for i := 0; i < max; i++ {

		//Abfrage ob einer der Strings schon zu ende ist
		if i > len(s1)-1 {
			result = result + string(s2[i])
		} else if i > len(s2)-1 {
			result = result + string(s1[i])
		} else {
			result = result + string(s1[i]) + string(s2[i])
		}

	}
	return result

}

func Zip(s1, s2 string) string {
	result := " "
	for i := 0; i < len(s1) || i < len(s2); i++ {
		if i < len(s1) {
			result += string(s1[i])
		} else if i < len(s2) {
			result += string(s2[i])
		}
	}
	return result
}
