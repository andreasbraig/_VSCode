package dllist_int

import "fmt"

// A doubly linked list.
// Contains a pointer to the anchor element.
type DLList struct {
	anchor *DLLNode
}

// NewDLList creates a new empty doubly linked list.
func NewDLList() *DLList {
	l := &DLList{}
	l.anchor = NewDLLNode()
	return l
}

// IsEmpty returns true if the list is empty.
func (l *DLList) IsEmpty() bool {
	return l.anchor.IsEmpty()
}

// PushBack adds a new element at the end of the list.
func (l *DLList) PushBack(value int) {
	l.anchor.Prev.Insert(value)
}

// PushFront adds a new element at the beginning of the list.
func (l *DLList) PushFront(value int) {
	// TODO
	l.anchor.Insert(value)
	// Hinweis: Sie können sich an der Methode PushBack() orientieren.
}

// PopBack removes the last element from the list.
// If the list is empty, nothing happens.
func (l *DLList) PopBack() {
	// TODO

	if l.IsEmpty() {
		return
	}

	l.anchor.Prev.Remove()

	// Hinweis: Sie können sich an der Methode PushBack() orientieren.
}

// PopFront removes the first element from the list.
// If the list is empty, nothing happens.
func (l *DLList) PopFront() {
	// TODO

	if l.IsEmpty() {
		return
	}

	l.anchor.Next.Remove()

	// Hinweis: Sie können sich an der Methode PushBack() orientieren.
}

//GetLength returns the length of the list.

// Get expects an index and returns the value at that index.
// If the index is out of bounds, it returns -1.
func (l *DLList) Get(index int) int {
	// TODO

	// Checke ob der index Falsch ist
	if index < 0 {
		return -1
	}
	// Einführung des Counters
	counter := 0

	// Hinweis: Verwenden Sie eine Schleife, die beim Anker beginnt und index mal
	for crt := l.anchor.Next; crt != l.anchor; crt = crt.Next {
		if counter == index {
			return crt.Value
		}
		counter++
	}

	// Hinweis: Schreiben Sie eine Schleife, beim Anker beginnend index mal
	// von Element zu Element springt.
	// Sie können Sich an der String()-Methode orientieren.
	return -1
}

// Insert inserts a new element at the given index with the given value.
func (l *DLList) Insert(idx, value int) {

	el := l.getpointer(idx)

	if el == nil {
		return
	}

	el.Insert(value)

}

func (l *DLList) InsertSorted(value int) { //Funktion nicht klar

	for crt := l.anchor.Next; crt != l.anchor; crt = crt.Next {
		if crt.Prev.Value < value && crt.Next.Value > value {
			crt.Insert(value)
		}
	}

}

// Remove removes the element at the given index.
func (l *DLList) Remove(idx int) {
	el := l.getpointer(idx)

	if el == nil {
		return
	}

	el.Remove()
}

// getpointer returns a pointer to the element at the given index.

// Replace expects an index and a value and replaces the value at that index.
func (l *DLList) Replace(idx, value int) {

	el := l.getpointer(idx)

	if el == nil {
		return
	}

	el.Value = value

}

// Swap expects two values, finds their corresponding elements and
// swaps these elements by rewriting their Pointers.
func (l *DLList) Swap(value1, value2 int) {
	// TODO

	idx1 := l.getindex(value1)

	idx2 := l.getindex(value2)

	e1 := l.getpointer(idx1)
	if e1 == nil {
		return
	}

	e2 := l.getpointer(idx2)
	if e2 == nil {
		return
	}

	swapNodes(e1, e2)

	// Hinweis: Verwenden Sie den gleichen Ansatz wie bei Get(), um die Elemente zu finden.
	// Verwenden Sie dann die Funktion swapDLLNodes(), um die Pointer zu tauschen.

	// Anmerkung: Es bietet sich an, eine Hilfsfunktion zu schreiben, die einen Index
	// erwartet und einen Pointer auf das entsprechende Element zurückgibt.
	// Diese Funktion können Sie dann sowohl in Get() als auch in Swap() verwenden.
}

// String returns a string representation of the list.
func (l *DLList) String() string {
	valuestrings := []string{}
	for e := l.anchor.Next; e != l.anchor; e = e.Next {
		valuestrings = append(valuestrings, e.String())
	}
	return fmt.Sprintf("%v", valuestrings)
}
