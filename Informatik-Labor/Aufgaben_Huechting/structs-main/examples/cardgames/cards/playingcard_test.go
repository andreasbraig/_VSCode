package cards

import "fmt"

func ExamplePlayingCard_String() {
	c1 := NewCard("Clubs", "Seven")
	c2 := NewCard("Spades", "Ace")

	fmt.Println(c1.String())
	fmt.Println(c1)
	fmt.Println(c2.String())
	fmt.Println(c2)

	// Output:
	// Seven of Clubs
	// Seven of Clubs
	// Ace of Spades
	// Ace of Spades
}
