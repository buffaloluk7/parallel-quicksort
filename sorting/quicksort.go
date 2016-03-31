package sorting

import (
	"sync"
	"sort"
)

func Quicksort(s sort.Interface, threshold int) {
	barrier := &sync.WaitGroup{}
	barrier.Add(1)

	go quick(s, 0, s.Len() - 1, barrier, threshold)

	barrier.Wait()
}

func quick(data sort.Interface, left, right int, barrier *sync.WaitGroup, threshold int) {
	defer barrier.Done()

	if left >= right {
		return
	}

	c := quicksortContext{data, left, right, left}

	pivotIndex := c.partition()

	pivotIsFirstElement := pivotIndex == 0
	if !pivotIsFirstElement {
		barrier.Add(1)

		sortParallel := threshold == 0 || (pivotIndex - 1) - left > threshold
		if sortParallel {
			go quick(data, left, pivotIndex - 1, barrier, threshold)
		} else {
			quick(data, left, pivotIndex - 1, barrier, threshold)
		}
	}

	pivotIsLastElement := pivotIndex + 1 == data.Len()
	if !pivotIsLastElement {
		barrier.Add(1)

		sortParallel := threshold == 0 || right - (pivotIndex + 1) > threshold
		if sortParallel {
			go quick(data, pivotIndex + 1, right, barrier, threshold)
		} else {
			quick(data, pivotIndex + 1, right, barrier, threshold)
		}
	}
}

type quicksortContext struct {
	data                 sort.Interface
	left, right, storage int
}

func (c quicksortContext) movePivotToEnd() {
	c.data.Swap(c.left, c.right)
}

func (c quicksortContext) movePivotInPlace() {
	c.data.Swap(c.storage, c.right)
}

func (c quicksortContext) pivotPosition() int {
	return c.storage
}

func (c quicksortContext) partition() int {
	c.movePivotToEnd()

	for i := c.left; i < c.right; i++ {
		if c.data.Less(i, c.right) {
			c.data.Swap(i, c.storage)
			c.storage += 1
		}
	}

	c.movePivotInPlace()

	return c.pivotPosition()
}