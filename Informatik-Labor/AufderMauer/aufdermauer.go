package aufdermauer

import "fmt"

func TanzeWanze() {

	stringvorlage := "Wanze"
	string1 := stringvorlage
	puffer := ""

	for i := len(string1); i > 0; i-- {

		fmt.Println(string1)

		puffer += string(string1[i-1])

		string1 = string1[0 : i-1]

	}

	fmt.Println(string1)

	for i := len(puffer); i > 0; i-- {

		string1 += string(puffer[i-1])

		fmt.Println(string1)

	}

}

func KillWord(s string) string {

	string1 := s

	if len(s) == 0 {
		return ""

	}
	if len(s) > 5 {
		fmt.Println(string1)
		return KillWord(s[0 : len(s)-2])
	}

	fmt.Println(string1)

	return KillWord(s[0 : len(s)-1])

}

// Erwartet ein Wort und buchstabiert dieses.
func BuildWord(s, result string) string {

	if len(result) == 0 {
		return ""
	}
	if len(s) == 4 && len(result) >= 2 {

		s += string(result)

		fmt.Println(s)

		return ""

	} else {

		head, tail := string(result[0]), string(result[1:])

		s += string(head)

		fmt.Println(s)

		return BuildWord(s, tail)

	}

}

func AufderMauer() {
	fmt.Println("Auf der Mauer auf der Lauer liegt ne kleine", KillWord("Wanze"))
}
