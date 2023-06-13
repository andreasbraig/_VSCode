package binsearchtree

//Wichtig: Ein Knoten ist leer, wenn beide kinder nil sind.

type Element struct {
	value int
	Left  *Element
	Right *Element
}

// NewElement erzeugt einen neuen leeren Knoten.
func NewElement() *Element {

	return &Element{}

}

// IsEmpty liefert true falls der Knoten leer ist
// Ein Knoten ist leer, wenn beide kinder Nil sind
func (e Element) IsEmpty() bool {
	return e.Left == nil && e.Right == nil
}

// SetValue erwartet einen Wert und schreibt ihn in den Knoten
// Erzeugt neue Kinder, falls der Knoten leer ist
func (e *Element) SetValue(value int) {

	e.value = value

	if e.IsEmpty() {
		e.Left = NewElement()
		e.Right = NewElement()
	}
}

// Insert fügt ein neues Element in den Baum ein.
// Kleiner geht nach Links größer gleich nach rechts!
func (e *Element) Insert(value int) {
	if e.IsEmpty() {
		e.SetValue(value)
		return
	}
	if value < e.value {
		e.Left.Insert(value)
		return
	} else {
		e.Right.Insert(value)

	}

}
