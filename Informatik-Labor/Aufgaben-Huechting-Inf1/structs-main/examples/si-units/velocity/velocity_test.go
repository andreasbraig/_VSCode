package velocity

import (
	"fmt"

	"github.com/tel22a-inf/go.structs/examples/si-units/distance"
	"github.com/tel22a-inf/go.structs/examples/si-units/duration"
)

// ExampleVelocity_constructors konstruiert verschiedene Velocity-Objekte und gibt sie aus.
func ExampleVelocity_constructors() {
	// Geschwindigkeit aus km/h-Angabe:
	v1 := FromKMH(42)
	fmt.Println(v1)

	// Geschwindigkeit aus m/s-Angabe:
	v2 := FromMPS(15)
	fmt.Println(v2)

	// Geschwindigkeit aus Meilen pro 2 Tage:
	distance := distance.FromMI(50)
	duration := duration.FromD(2)
	v3 := NewVelocity(distance, duration)
	fmt.Println(v3)

	// Output:
	// {42000000 3600000}
	// {15000 1000}
	// {80467000 172800000}
}
