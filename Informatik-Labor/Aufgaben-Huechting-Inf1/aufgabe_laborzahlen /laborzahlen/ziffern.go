package laborzahlen

func Digits(n int) []int {
	if n == 0 {
		return []int{}
	}
	return append(Digits(n/10), n%10)
}

func ToString(n int) string {

	result := ""

	for _, element := range Digits(n){

		result += string('0' + byte(element))

	}

	return result
}
