package main

import (
	"fmt"
)

func import_Int(msg string) int {

	var input int

	fmt.Print(msg)
	fmt.Scanln(&input)
	fmt.Println("Sie haben", input, "eingegeben.")

	return input
}


func import_String(msg string) string {

	var input string

	fmt.Println(msg)
	fmt.Scanln(&input)
	fmt.Println("Sie haben", input, "eingegeben ")


	return input 

}
