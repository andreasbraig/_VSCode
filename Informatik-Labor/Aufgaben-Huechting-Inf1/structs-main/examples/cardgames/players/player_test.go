package players

import (
	"fmt"

	"github.com/tel22a-inf/go.structs/examples/cardgames/cards"
)

func ExampleNew() {
	p1 := New("Max Mustermann")
	fmt.Println(p1)
	fmt.Println()

	p2 := New(
		"Monika Musterfrau",
		cards.NewCard("hearts", "seven"),
		cards.NewCard("spades", "ten"),
	)
	fmt.Println(p2)
	fmt.Println()

	// Output:
	// Max Mustermann: Keine Handkarten
	//
	// Monika Musterfrau:
	// seven of hearts
	// ten of spades
}
