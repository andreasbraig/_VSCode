package selectionsort

import (
	"testing"

	"github.com/tel22a-inf/sorting/testhelpers"
)

func TestSelectionSort(t *testing.T) {
	list := []int{3, 2, 1}
	SelectionSort(list)
	testhelpers.AssertListsEqual(t, []int{1, 2, 3}, list)
}
