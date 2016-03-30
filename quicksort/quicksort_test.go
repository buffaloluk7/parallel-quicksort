package quicksort

import (
	"testing"
	"sort"
"math/rand"
)

func TestQuicksort(t *testing.T) {

	numberOfItems := 1000
	items := make([]int, numberOfItems)
	for i, _ := range items {
		items[i] = rand.Int() % 200
	}

	s := sort.IntSlice(items)

	Quick(s)

	if !sort.IsSorted(s) {
		t.Fail()
	}
}