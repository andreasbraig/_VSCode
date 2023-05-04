package velocity

import (
	"github.com/tel22a-inf/go.structs/examples/si-units/distance"
	"github.com/tel22a-inf/go.structs/examples/si-units/duration"
)

// Velocity ist ein Typ, der Geschwindigkeiten speichert.
// Eine Geschwindigkeit ist ein Quotient aus Entfernung und Zeit,
// daher werden hier Distance und Duration verwendet.
type Velocity struct {
	dist distance.Distance
	dur  duration.Duration
}

// NewVelocity erwartet eine Distance und eine Duration und
// konstruiert daraus ein Velocity-Objekt.
func NewVelocity(dist distance.Distance, dur duration.Duration) Velocity {
	return Velocity{dist: dist, dur: dur}
}

// FromKMH liefert ein Velocity-Objekt zur gegebenen KM-Pro-Stunde-Angabe.
func FromKMH(kmh int) Velocity {
	return NewVelocity(distance.FromKM(kmh), duration.FromH(1))
}

// FromMPS liefert ein Velocity-Objekt zur gegebenen Meter-Pro-Sekunde-Angabe.
func FromMPS(ms int) Velocity {
	return NewVelocity(distance.FromM(ms), duration.FromS(1))
}
