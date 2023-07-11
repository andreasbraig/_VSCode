package main

import "fmt"

type Person struct {
	Name           string
	Age            int
	Benkanntschaft *Person
}

func main() {

	Andi := Person{Name: "Andi", Age: 22}
	Marsch := Person{Name: "Marsch", Age: 65, Benkanntschaft: &Andi}

	fmt.Println(Marsch)

}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s\nAge: %v\nBekanntschaft:%s", p.Name, p.Age, p.Benkanntschaft.Name)
}


