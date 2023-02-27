package duration

// Duration ist ein Datentyp für Zeitdauern.
// Als Basis-Einheit wird eine Zeitdauer in Millisekunden gespeichert.
type Duration int

// FromMS konstruiert eine Duration aus einer Millisekunden-Angabe.
func FromMS(ms int) Duration {
	return Duration(ms)
}

// FromS konstruiert eine Duration aus einer Sekunden-Angabe.
func FromS(s int) Duration {
	return FromMS(s * 1000)
}

// FromM konstruiert eine Duration aus einer Minuten-Angabe.
func FromM(m int) Duration {
	return FromS(m * 60)
}

// FromH konstruiert eine Duration aus einer Stunden-Angabe.
func FromH(h int) Duration {
	return FromM(h * 60)
}

// FromD konstruiert eine Duration aus einer Tages-Angabe.
func FromD(d int) Duration {
	return FromH(d * 24)
}
