package main

import (
	"fmt"
)

func Devides(m,n int) bool {

	var result int 

	result = n / m 

	if m * result == n{
		return true
	}
	return false

}

func Modulo(m,n int) bool {
	result := n % m

	if result > 0 {
		return false
	}
	return true
}

func main(){

	fmt.Println(Devides(5,20))
	fmt.Println(Modulo(5,20))

}