package aufgabe6

import "fmt"

func ExampleCommonFriends() {
	f1 := Person{"Gödel", []Person{}}
	f2 := Person{"Turing", []Person{}}
	f3 := Person{"Euler", []Person{}}
	f4 := Person{"Gauss", []Person{}}
	f5 := Person{"Lovelace", []Person{}}

	p1 := Person{"Lamport", []Person{f1, f3, f5}}
	p2 := Person{"Knuth", []Person{f2, f3, f4}}

	fmt.Println(CommonFriends(p1, p2))
	fmt.Println(CommonFriends(p1, p1))
	fmt.Println(CommonFriends(p2, p2))

	// Output:
	// [{Euler []}]
	// [{Gödel []} {Euler []} {Lovelace []}]
	// [{Turing []} {Euler []} {Gauss []}]
}
