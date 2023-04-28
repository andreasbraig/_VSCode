package linkedlist

type DLElement struct {
	Prev  *DLElement
	Value int
	Next  *DLElement
}

// NewDLElement creates a new empty doubly liked list Element
func NewDLElement() *DLElement {
	e := &DLElement{}
	e.Prev = e
	e.Next = e

	return e

}


// IsEmpty returns true if the element is empty
func (e *DLElement) IsEmpty() bool {

	return e.Prev == e && e.Next == e

}

// Inserts a new element after the current element
func (e *DLElement) Insert(value int) *DLElement{
	
	newElement := &DLElement{Prev: e,Value: value, Next: e.Next}
	e.Next.Prev = newElement
	e.Next = newElement
	return newElement

}
