package mergesort

import (
	"testing"

	"github.com/tel22a-inf/sorting/testhelpers"
)

func TestMerge(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9}
	list2 := []int{2, 4, 6, 8, 10}

	list := Merge(list1, list2)
	testhelpers.AssertListsEqual(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, list)
}
