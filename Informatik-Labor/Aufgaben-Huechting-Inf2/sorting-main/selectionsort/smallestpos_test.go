package selectionsort

import "testing"

func TestSmallestPos(t *testing.T) {
	list := []int{3, 2, 1}
	pos := SmallestPos(list)
	if pos != 2 {
		t.Errorf("SmallestPos(%v) = %v, want %v", []int{3, 2, 1}, pos, 2)
	}
}
