package quicksort

import (
	"sync"
	"sort"
)

func Quick(s sort.Interface) {
	barrier := &sync.WaitGroup{}
	barrier.Add(1)
	go quick(s, 0, s.Len() - 1, barrier)
	barrier.Wait()
}

func quick(data sort.Interface, left, right int, barrier *sync.WaitGroup) {
	defer barrier.Done()

	if left >= right {
		return
	}

	c := quicksortContext{data, left, right, left}

	pivotIndex := c.partition()

	pivotIsFirstElement := pivotIndex == 0
	if !pivotIsFirstElement {
		barrier.Add(1)
		go quick(data, left, pivotIndex - 1, barrier)
	}

	pivotIsLastElement := pivotIndex + 1 == data.Len()
	if !pivotIsLastElement {
		barrier.Add(1)
		go quick(data, pivotIndex + 1, right, barrier)
	}
}

type quicksortContext struct {
	data sort.Interface
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