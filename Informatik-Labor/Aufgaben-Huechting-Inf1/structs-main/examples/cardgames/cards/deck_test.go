package cards

import "fmt"

func ExampleDeck_String() {
	d1 := NewDeck(
		NewCard("Spades", "Ace"),
		NewCard("Spades", "Queen"),
		NewCard("Hearts", "Queen"),
		NewCard("Hearts", "Seven"),
	)

	fmt.Println(d1)

	// Output:
	// Ace of Spades
	// Queen of Spades
	// Queen of Hearts
	// Seven of Hearts
}

func ExampleDeck_Length() {
	d1 := NewDeck(
		NewCard("Spades", "Ace"),
		NewCard("Spades", "Queen"),
		NewCard("Hearts", "Queen"),
		NewCard("Hearts", "Seven"),
	)

	fmt.Println(d1.Length())

	// Output:
	// 4
}

func ExampleDeck_Get() {
	d1 := NewDeck(
		NewCard("Hearts", "Seven"),
		NewCard("Spades", "Ace"),
		NewCard("Spades", "King"),
	)

	fmt.Println(d1.Get(0))
	fmt.Println(d1.Get(1))
	fmt.Println(d1.Get(2))

	// Output:
	// Seven of Hearts
	// Ace of Spades
	// King of Spades
}

func ExampleDeck_Add() {
	d1 := NewDeck()
	d1.Add(
		NewCard("Hearts", "Seven"),
		NewCard("Spades", "Ace"),
		NewCard("Spades", "King"),
	)

	fmt.Println(d1.Get(0))
	fmt.Println(d1.Get(1))
	fmt.Println(d1.Get(2))

	// Output:
	// Seven of Hearts
	// Ace of Spades
	// King of Spades
}

func ExampleDeck_Remove() {
	d1 := NewDeck(
		NewCard("Hearts", "Seven"),
		NewCard("Spades", "Ace"),
		NewCard("Spades", "King"),
		NewCard("Clubs", "Eight"),
		NewCard("Diamonds", "Jack"),
	)

	d1.Remove(0)
	fmt.Println(d1)

	d1.Remove(2)
	fmt.Println(d1)

	d1.Remove(2)
	fmt.Println(d1)

	// Output:
	// Ace of Spades
	// King of Spades
	// Eight of Clubs
	// Jack of Diamonds
	//
	// Ace of Spades
	// King of Spades
	// Jack of Diamonds
	//
	// Ace of Spades
	// King of Spades
}

func ExampleDeck_DrawCard() {
	d1 := NewDeck(
		NewCard("Hearts", "Seven"),
		NewCard("Spades", "Ace"),
		NewCard("Spades", "King"),
	)

	c1 := d1.DrawCard()
	fmt.Println(c1)
	fmt.Println(d1)

	// Output:
	// King of Spades
	// Seven of Hearts
	// Ace of Spades
}

func ExampleDeck_Shuffle() {
	d1 := NewDeck(
		NewCard("Hearts", "Seven"),
		NewCard("Spades", "Ace"),
		NewCard("Spades", "King"),
	)

	d1.Shuffle()
	fmt.Println(d1)

	// Hinweis: d1.Shuffle() liefert immer das selbe Ergebnis, weil wir in diesem
	// Test keinen Random-Seed vergeben haben. Dadurch wird vom Zufallsgenerator immer
	// die selbe Zufallsfolge erzeugt. Dies ist sehr praktisch für die Tests, da wir
	// auf diese Weise prüfen können, ob Shuffle() arbeitet, und trotzdem das erwartete
	// Ergebnis kennen.
	// In einem echten Programm darf man nicht vergessen, z.B. die folgende Zeile
	// z.B. zu Beginn der main()-Funktion auszuführen:
	// rand.Seed(time.Now().UnixNano())

	// Output:
	// Seven of Hearts
	// King of Spades
	// Ace of Spades
}

func ExampleNewDeck32() {
	d1 := NewDeck32()

	fmt.Println(d1)

	// Output:
	// Seven of Clubs
	// Eight of Clubs
	// Nine of Clubs
	// Ten of Clubs
	// Jack of Clubs
	// Queen of Clubs
	// King of Clubs
	// Ace of Clubs
	// Seven of Spades
	// Eight of Spades
	// Nine of Spades
	// Ten of Spades
	// Jack of Spades
	// Queen of Spades
	// King of Spades
	// Ace of Spades
	// Seven of Hearts
	// Eight of Hearts
	// Nine of Hearts
	// Ten of Hearts
	// Jack of Hearts
	// Queen of Hearts
	// King of Hearts
	// Ace of Hearts
	// Seven of Diamonds
	// Eight of Diamonds
	// Nine of Diamonds
	// Ten of Diamonds
	// Jack of Diamonds
	// Queen of Diamonds
	// King of Diamonds
	// Ace of Diamonds
}
