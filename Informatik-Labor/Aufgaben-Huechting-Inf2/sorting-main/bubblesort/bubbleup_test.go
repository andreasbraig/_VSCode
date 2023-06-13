package bubblesort

import (
	"fmt"
	"testing"

	"github.com/tel22a-inf/sorting/testhelpers"
)

func TestBubbleUp(t *testing.T) {
	list := []int{1, 4, 3, 2, 6, 5}

	swapped := BubbleUp(list)
	testhelpers.AssertTrue(t, swapped, fmt.Sprintf("BubbleUp(%v)", []int{1, 4, 3, 2, 6, 5}))
	testhelpers.AssertListsEqual(t, []int{1, 3, 2, 4, 5, 6}, list)

	swapped = BubbleUp(list)
	testhelpers.AssertTrue(t, swapped, fmt.Sprintf("BubbleUp(%v)", []int{1, 3, 2, 4, 5, 6}))
	testhelpers.AssertListsEqual(t, []int{1, 2, 3, 4, 5, 6}, list)

	swapped = BubbleUp(list)
	testhelpers.AssertFalse(t, swapped, fmt.Sprintf("BubbleUp(%v)", []int{1, 2, 3, 4, 5, 6}))
}
