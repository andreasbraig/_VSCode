package intro

import "fmt"

type Kunde struct {
	Name    string
	Vorname string
	Kdnr    string
	Punkte  int
}

// Zeigt wie man einen Struct einführt und definiert über den var befehl
func Demo1() {

	//Einführung der Variable
	var k1 Kunde

	//Ausgabe des leeren Structs
	fmt.Println(k1)

	//deklarieren der Kundendaten
	k1.Name = "Mustermann"
	k1.Vorname = "Max"
	k1.Kdnr = "A40991"
	k1.Punkte = 42

	//Ausgabe des Kunden
	fmt.Println(k1)

	k2 := Kunde{
		Name:    "Beispiel",
		Vorname: "Barbara",
		Kdnr:    "465b",
		Punkte:  42,
	}

	fmt.Println(k2)

}

// Liefert eine repräsentation des Structs Kunde
/*
func (kunde Kunde) String() string {

	return fmt.Sprintf("KUNDE:\n Name: %v\n Vorname: %v\n Kundennummer: %v\n Punkte: %v\n", kunde.Name, kunde.Vorname, kunde.Kdnr, kunde.Punkte)

}
*/
