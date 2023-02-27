package formen

import "fmt"

type Quadrat struct {
	A int
}

type Rechteck struct {
	A int
	B int
}

func (q Quadrat) Draw() {

	for i := 0; i < q.A; i++ {

		for j := 0; j < q.A; j++ {

			fmt.Print("*")

		}

		fmt.Println()

	}

}

func (r Rechteck) Draw() {

	for i := 0; i < r.A; i++ {

		for j := 0; j < r.B; j++ {

			fmt.Print("*")

		}

		fmt.Println()

	}

}
