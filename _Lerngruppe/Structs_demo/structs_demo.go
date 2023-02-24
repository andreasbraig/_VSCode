package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {

	Simon := person{"Simon",20}

	Hannah := person{"Hannah",19}

	fmt.Println(Simon)
	fmt.Println(Hannah)

	Hannah.age = 23

	fmt.Println(Hannah)

	Lucas := Simon

	Lucas.name = "Lucas"

	fmt.Println(Lucas)

	fmt.Println(Lucas.age)


	

}



