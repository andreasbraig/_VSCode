package circularbuffer

// Circularbuffer is a struct that holds the buffer, size, and
// positions of first and last elements of the buffer.
type Circularbuffer struct {
	buffer []interface{}
	size   int
	first  int
}

// New creates a new Circularbuffer with the given size.
func New(capacity int) *Circularbuffer {
	return &Circularbuffer{
		buffer: make([]interface{}, capacity),
		size:   0,
		first:  0,
	}
}

// Size returns the number of elements in the buffer.
func (cb *Circularbuffer) Size() int {
	return cb.size
}

// Capacity returns the capacity of the buffer.
func (cb *Circularbuffer) Capacity() int {
	return len(cb.buffer)
}

// IsEmpty returns true if the buffer is empty.
func (cb *Circularbuffer) IsEmpty() bool {
	return cb.size == 0
}

// IsFull returns true if the buffer is full.
func (cb *Circularbuffer) IsFull() bool {
	return cb.size == len(cb.buffer)
}

// Add adds an element to the buffer.
// If the buffer is full, the oldest element is overwritten.
func (cb *Circularbuffer) Add(element interface{}) {
	// TODO
	if cb.IsFull() {
		cb.RemoveOldest() // First eins nach vorne setzen
	}

	last := (cb.first + cb.size) % cb.Capacity() // Last ist die first position plus die size 

	cb.buffer[last] = element // ändere den buffer an stelle last 

	cb.size++ // erhöhe size 
}

// GetAndRemove returns the oldest element in the buffer and removes it.
// If the buffer is empty, a panic occurs.
func (cb *Circularbuffer) GetAndRemove() interface{} {
	// TODO

	el := cb.GetOldest()

	cb.RemoveOldest()

	return el
}

// GetOldest returns the oldest element in the buffer.
// If the buffer is empty, a panic occurs.
func (cb *Circularbuffer) GetOldest() interface{} {
	// TODO
	if cb.IsEmpty() {
		panic("this circular buffer is empty")
	}

	return cb.buffer[cb.first]

}

// RemoveOldest removes the oldest element in the buffer.
// If the buffer is empty, a panic occurs.
func (cb *Circularbuffer) RemoveOldest() {
	// TODO
	if cb.IsEmpty() {
		panic("Buffer is empty")
	}
	cb.first = (cb.first + 1) % cb.Capacity() // First überprüfen und reduzieren
	cb.size-- // size reduzieren 
}
