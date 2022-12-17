package main

import (
	"fmt"
)

func checkDivides(m, n int) bool {

	if m == 0{
		return false
	}

	var result int

	result = n / m
	if m*result == n {
		return true
	}
	return false

}

func Modulo(m, n int) bool {
	result := n % m

	if result == 0 {
		return true
	}
	return false
}

func main() {

	fmt.Println(checkDivides(3, 15))
	fmt.Println(checkDivides(3, 10))
	fmt.Println(checkDivides(2, 10))
	fmt.Println(checkDivides(7, 21))
	fmt.Println(checkDivides(15, 60))
	fmt.Println(checkDivides(-3, 60))
	fmt.Println(checkDivides(-4, 8))
	fmt.Println(checkDivides(-4, 9))
	fmt.Println(checkDivides(10, 3))
	fmt.Println(checkDivides(0, 4))
	/*fmt.Println(checkDivides(10, -4))
	fmt.Println(checkDivides(25, -2))
	fmt.Println(checkDivides(25, -5))*/
}
