package sorting

import (
	"testing"
	"math/rand"
	"log"
	"time"
	"sort"
)

var quicksortTestData = []struct {
	input     []int
	threshold int
}{
	//Simple Tests
	{[]int{}, -1},
	{[]int{42, 45, 78, 9, 5, 0, 0, 1, 1}, 0},
	{[]int{42, 23}, 0},

	//Tests with 1 million elements
	{rand.Perm(1000000), 0},
	{rand.Perm(1000000), 500000},
	{rand.Perm(1000000), 100000},
	{rand.Perm(1000000), 10000},
	{rand.Perm(1000000), 1000},
	{rand.Perm(1000000), 100},
	{rand.Perm(1000000), 10},

	//Tests with 100 million elements
	{rand.Perm(100000000), 0},
	{rand.Perm(100000000), 80000000},
	{rand.Perm(100000000), 50000000},
	{rand.Perm(100000000), 10000000},
}

func TestQuicksort(t *testing.T) {
	for _, testCase := range quicksortTestData {
		start := time.Now()

		input := sort.IntSlice(testCase.input)
		Quicksort(input, testCase.threshold)

		if (!sort.IsSorted(input)) {
			t.Fail()
		}

		log.Printf("Took %s to sort %d items (with threshold %d).", time.Since(start), len(testCase.input), testCase.threshold)
	}
}