package arztpraxis

import "fmt"



func Example() {
	m1 := Medikament{
		Name: "Ibuprofen",
		Unvertraeglichkeiten: []Medikament{},
	}
	m2 := Medikament{
		Name: "Paracetamol",
		Unvertraeglichkeiten: []Medikament{m1},
	}
	/*m3 := Medikament{
		Name: "Nasenspray",
		Unvertraeglichkeiten: []Medikament{},
	}*/

	m1.AddUnvertraeglichkeit(m2)

	p1 := Patient{
		Name: "Donald Knuth",
		Medis: []Medikament{m1},
		Unvis: []Medikament{},
	}

	fmt.Println(p1.CheckMedi(m2))


	//Output:
	//false



}