package strings

// Erwartet einen String s ind liefert
// einen String, bei dem alle Zeichen
// aus s klein geschrieben sind.
func ToLower(s string) string {
	// TODO
	result := ""

	for _, character := range s {

		if character >= 'A' && character <= 'Z' { //Wenn der Charakter zwischen ASCII GroßA und Groß Z Liegt
			character += ('a' - 'A') // Addiere den Zahlenbereich in ASCII zwischen klein A und groß A, dass der Buchstabe Klein wird
		}

		if character >= 'Ä' && character <= 'Ü' { //Wenn der Charakter zwischen ASCII GroßA und Groß Z Liegt
			character += ('a' - 'A') // Addiere den Zahlenbereich in ASCII zwischen klein A und groß A, dass der Buchstabe Klein wird
		}

		result += string(character) //String wieder zusammenbauen

	}

	return result

	// Alternative:
	// return strings.ToLower(s)

}
