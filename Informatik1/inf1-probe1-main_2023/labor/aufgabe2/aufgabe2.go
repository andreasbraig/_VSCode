package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Colour.
PUNKTE: 10
BEWERTUNG:
*/

// Colour erwartet eine Zahl t, die als Sekundenangabe aufgefasst werden soll.
// Die Funktion liefert einen String, der angibt, welche Farbe eine Ampel nach t Sekunden hat.
// Dabei soll die Ampel bei Rot starten und alle 10 Sekunden umschalten.
func Colour(t int) string {
	// TODO

	x := t % 40 
	
	if x < 0 {
		x += 40
	}
	

	// bei 0-10 Ampel rot 
	// bei 11-20 Ampel gelb
	// bei 21-30 Ampel grün 
	// restart 

	if x >= 0 && x < 10 {
		return "Rot"
	}else if x >= 10 && x < 20 {
		return "Rot-Gelb"
	}else if x >= 20 && x < 30 {
		return "Grün"
	}else if x >= 30 && x < 40 {
		return "Gelb"
	}


	return ""
}
