package formen

import "fmt"

func ExampleQuadrat_Draw() {

	q1 := Quadrat{A: 4}
	r1 := Rechteck{A: 4, B: 15}

	fmt.Println(q1.A)
	q1.Draw()
	fmt.Println()
	fmt.Println(r1.A, r1.B)
	r1.Draw()

	q1.Resize(2)
	r1.Resize(2)

	fmt.Println(q1.A)
	q1.Draw()
	fmt.Println()
	fmt.Println(r1.A, r1.B)
	r1.Draw()

	//Output:

}
