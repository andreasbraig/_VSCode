package foo

// Erzeigt eine leere Liste von Strings
func CreateEmptyStringSlice() []string {
	return []string{}
}

//erzeugt eine Liste vo Strings mit einem Anfangswert.

func CreateStringSliceWithOneElement(start string) []string {

	return []string{start}
}

//Durchsucht die Liste nach dem Wort W und gibt die Position des ersten Vorkommens zurück

func FindString(list []string, w string) int {

	for idx, element := range list {
		if element == w {
			return idx
		}
	}

	return -1

}

// Erwartet zwei listen und ein wort w. Sucht die Position von w in list 1 und liefert das entsprechende Wort in list 2

func LookupString(list1, list2 []string, w string) string {

	pos := FindString(list1, w)
	if pos == -1 || pos >= len(list2) {
		return " "
	}
	return list2[pos]

}
