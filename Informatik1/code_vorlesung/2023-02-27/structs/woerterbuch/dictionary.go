package woerterbuch

type Dictionary struct {
	entries []Eintrag
}

// Erwartet ein Deutsches Wort
// Liefert das jeweilige Englische Wort aus dem Wörterbuch
func (d Dictionary) LookupEn(de string) string {

	for _, eintrag := range d.entries {
		if eintrag.De == de {
			return eintrag.En
		}

	}
	return ""
}

// Erwartet ein Deutsches und Englisches wort
// Erweitert das Wörterbuch um den Eintrag aus den beiden Wörtern
func (d *Dictionary) Add(de, en string) {

	newEntry := Eintrag{de, en}

	//Aus dem Pointer bitte den Wert machen, daran appenden und wieder in die Ursprungsliste schreiben.
	d.entries = append(d.entries, newEntry)

}
