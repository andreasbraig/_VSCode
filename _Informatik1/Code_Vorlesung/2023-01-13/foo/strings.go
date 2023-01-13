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


	for idx,element := range list {
		if element == w {
			return idx
		}
	}

	return 0

}
