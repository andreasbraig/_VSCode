package main

import (
	"fmt"
)

func Listen() {
	l1 := []int{13, 2, 4, 25, 4, 12, 2, -3, 15, 25, 13, 4, 75, 42, 2, 4, 13, 10, 4, 17, 38}

	n := 2

	result := 0

	count := 0

	for _, element := range l1 {
		if element > n {
			result++
		}

		count++

		fmt.Println(element)

		
	}

	fmt.Println(len(l1))

	fmt.Println(count, result)
}

func Calc() bool{
	var n int = 13 
	var m int = 13

	return n < m 
}

func main() {

	l1 := []int{1, 4, 5, 6, 7, 9, 5}

	fmt.Println(ContainsChain(l1))

}


func ContainsChain(list []int) bool {
	// TODO

	x := 0

	for _, element := range list[:len(list)-1]  {

		fmt.Println(element)
		fmt.Println(element + 1)

		if element == x && element + 1 == x + 1 {
			return true 
		}

		x++

	}

	return false
}