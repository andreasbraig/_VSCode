package main

import "fmt"

type person struct {
	name string
	age  int
}

type Personenliste []person

func main() {

	namelist := []string{"Lars", "Silas", "Hannah", "Simon", "Lucas", "Andi"}
	agelist := []int{19, 23, 23, 20, 21, 21}

	Personlist := Personenliste{}

	for i, e := range namelist { //Führt jeden Namen mit einem alter in einer Person zusammen. Die Personen werden in eine Persinenliste zusammengesetzt

		Personlist = append(Personlist, Persongenerator(e, agelist[i]))

	}

	fmt.Println(Personlist)

	fmt.Println(person{"Lucas Weyland", 23})

}

// Erwartet einen Namen und ein Alter
// Führt diese im Datentyp Person zusammen
func Persongenerator(name string, age int) person {

	return person{name, age}

}

// Liefert eine Repräsentation des Datentyps Person
func (Person person) String() string {

	result := ""

	result += fmt.Sprintf("[%v, %v]", Person.name, Person.age)

	return result

}

// Lieferteine Representation des Dartentyps Personenliste
func (beispielliste Personenliste) String() string {

	result := ""

	for _, Row := range beispielliste {

		result += fmt.Sprintf("%v \n", Row)

	}

	return result

}
