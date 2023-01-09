package main

import (
	"fmt"
	"strings"
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

	fmt.Println(strings.ToLower("AHABDB"))


}
