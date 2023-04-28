package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ContainsPrefix.
PUNKTE: 10
BEWERTUNG:
HINWEIS: Die Funktion HasPrefix aus dem Package strings könnte hilfreich sein.
         Zu Übungszwecken können Sie diese Aufgabe sowohl mit als auch ohne
		 diese Hilfsfunktion lösen.
*/

// ContainsPrefix erwartet eine Liste von Strings sowie einen String s.
// Die Funktion liefert true, falls einer der Strings aus der Liste mit s beginnt.
func ContainsPrefix(list []string, s string) bool {
	// TODO

	for _, e := range list {

		if CheckPrefix(s, e) {
			return true
		}

	}

	return false
}

func CheckPrefix(Prefix, s string) bool {

	l := len(Prefix)

	return s[:l] == Prefix

}
