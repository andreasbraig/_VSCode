package game

import (
	"fmt"

	"github.com/tel22a-inf/go.structs/examples/cardgames/cards"
	"github.com/tel22a-inf/go.structs/examples/cardgames/players"
)

func ExampleDeal() {
	d1 := cards.NewDeck(
		cards.NewCard("Hearts", "Seven"),
		cards.NewCard("Spades", "Ace"),
		cards.NewCard("Spades", "King"),
		cards.NewCard("Diamonds", "Ten"),
	)
	d2 := players.New("Player1")

	Deal(&d1, &d2, 2)
	fmt.Println(d1)
	fmt.Println(d2)

	Deal(&d1, &d2, 1)
	fmt.Println(d1)
	fmt.Println(d2)

	// Output:
	// Seven of Hearts
	// Ace of Spades
	//
	// Player1:
	// Ten of Diamonds
	// King of Spades
	//
	// Seven of Hearts
	//
	// Player1:
	// Ten of Diamonds
	// King of Spades
	// Ace of Spades
}
