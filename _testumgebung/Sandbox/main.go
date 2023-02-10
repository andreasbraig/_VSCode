package main

import "fmt"

func main() {

	fmt.Println(ContainsSwitchedLetters("ABC", "BCA"))

}



func ContainsSwitchedLetters(s1, s2 string) bool {
	// TODO

	if len(s1) != len(s2){
		return false
	}

	for i := 0;i < len (s1)-1; i++ {
		if s1[i]!= s2[i]{
			if s2 == SwitchLetters(s1,i,i+1) {
				return true
			}
		}

	}

	

	return false
}

func SwitchLetters(s string, i, j int) string {
	result := []rune(s)
	// TODO

	result[i], result[j] = result[j], result[i]

	return string(result)
}