package main

import "fmt"

func Foo(n int) {
	if n == 0 { //Rekursionsanker (Abbruchbedingung)
		return
	}
	fmt.Println(n)
	Foo(n - 1)
	fmt.Println(n)
}

func RecursiveAddition(m, n int) int {
	if n == 0 {
		return m
	}
	return 1 + RecursiveAddition(m, n-1)
}

func main() {

	fmt.Println(2, 3)

}
