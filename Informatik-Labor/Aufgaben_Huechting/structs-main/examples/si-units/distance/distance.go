package distance

// Distance ist ein Datentyp für Entfernungen.
// Als Basis-Einheit wird eine Entfernung in Millimetern gespeichert.
type Distance int

// FromMM konstruiert eine Entfernung aus einer Millimeter-Angabe.
func FromMM(mm int) Distance {
	return Distance(mm)
}

// FromCM konstruiert eine Entfernung aus einer Zentimeter-Angabe.
func FromCM(cm int) Distance {
	return FromMM(cm * 10)
}

// FromM konstruiert eine Entfernung aus einer Meter-Angabe.
func FromM(m int) Distance {
	return FromCM(m * 100)
}

// FromKM konstruiert eine Entfernung aus einer Kilometer-Angabe.
func FromKM(km int) Distance {
	return FromM(km * 1000)
}

// FromMI konstruiert eine Entfernung aus einer Meilen-Angabe.
func FromMI(mi int) Distance {
	return FromCM(mi * 160934)
}
