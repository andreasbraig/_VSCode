package woerterbuch

import "fmt"

func ExampleDictionary_LookupDe() {

	d1 := Dictionary{}

	d1.Add("Haus", "house")
	d1.Add("Fahrrad", "bike")
	d1.Add("Fahrrad", "bycicle")

	d1.Add("Baum", "tree")

	fmt.Println(d1)

	fmt.Println(d1.LookupEn("Haus"))
	fmt.Println(d1.LookupEn("Fahrrad"))
	fmt.Println(d1.LookupEn("Baum"))

	//Output:

}
