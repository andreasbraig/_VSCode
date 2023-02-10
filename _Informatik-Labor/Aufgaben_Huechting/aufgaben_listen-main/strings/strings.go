package stringfuncs

import "strings"

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {
	// ANMERKUNG: Diese Funktion ist ein Beispiel, hier ist (noch) nichts zu tun.
	return CountChar(s,'A')
}

// Erwartet einen string s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {
	result := 0
	// TODO: Implementieren Sie die Funktion.
	// Ändern Sie anschließend CountA(),
	// so dass insgesamt möglichst wenig doppelter Code existiert.

	for _, char := range s {
		if char == c {
			result++
		}
	}

	return result
}

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	// ANMERKUNG: Diese Funktion ist ein Beispiel, hier ist nichts zu tun.
	result := ""
	for _, char := range s {
		result += string(char)
		result += string(char)
	}
	return result
}

// Erwartet einen String s und liefert s rückwärts zurück.
func Reverse(s string) string {
	result := ""
	// TODO

	for _, char := range s {

		result = string(char) + result // funktioniert wie += nur eher als =+ also erst neuer buchstabe dann rest, es wird umgekehrt angereiht

	  }

	return result
}

// Erwartet zwei Strings s1 und s2 und prüft, ob der eine der andere umgedreht ist.
func IsReverse(s1, s2 string) bool {
	// TODO

	/*
	
	eine mögliche Lösung: aber einfacher über Funktionen


	if len(s1) != len(s2) { // fängt nicht gleichlange Strings ab 
		return false
	}

	s1inv := ""

	for _, char := range s1 {

		s1inv = string(char) + s1inv

	}

	result := 0

	for i := range s1inv {      //Für die Länge "Range" der Kürzeren Liste werden beide Listen verglichen
		if s1inv[i] == s2[i] {
			result++
		}
	}

	return result == len(s2)*/

	return s1 == Reverse(s2)

}

// Erwartet einen String s und prüft, ob s ein Palindrom ist.
// Ein Palindrom ist ein String, der vorwärts und rückwärts gelesen gleich ist.
func IsPalindrome(s string) bool {
	// TODO
	return s == Reverse(s)

}
// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	// HINWEIS: Zählen Sie für jeden Buchstaben in s1 und s2, wie oft er vorkommt.

	for _, element := range s1 {

	if CountChar(s1, element) != CountChar(s2, element) {
		return false
	}

	}


	return true
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	// HINWEIS: Verwenden Sie die Funktion strings.ToLower() oder strings.ToUpper()

	s1low := strings.ToLower(s1)
	s2low := strings.ToLower(s2)

	for _, element := range s1low {

		

		if CountChar(s1low, element) != CountChar(s2low, element) {
			return false
		}
	
		}


	return true
}
