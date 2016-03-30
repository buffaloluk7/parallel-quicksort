package parallel_quicksort

import (
	"sync"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func SortQuick(items []int) {
	quicksort(0, len(items) - 1, items)
}

func quicksort(left, right int, items []int) {
	if left < right {
		pivotIndex := partition(left, right, items)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer func() { wg.Done() }()
			quicksort(left, pivotIndex - 1, items)
		}()
		go func() {
			defer func() { wg.Done() }()
			quicksort(pivotIndex + 1, right, items)
		}()

		wg.Wait()
	}
}

func partition(left, right int, items []int) int {
	i := left
	j := right - 1
	pivot := items[right]
	log.Debugf("Left: %d, right: %d, pivot value: %d (index: %d)", i, j, pivot, right)

	for {

		iChan := nextI(i, right, pivot, items)
		jChan := nextJ(j, left, pivot, items)

		i = <-iChan
		j = <-jChan

		if i < j {
			log.Debugf("Swap values")
			items[i], items[j] = items[j], items[i]
		} else {
			break
		}
	}

	if items[i] > pivot {
		items[i], items[right] = items[right], items[i]
	}

	return i
}

func nextI(i, right, pivot int, items []int) <-chan int {

	nextIChan := make(chan int)

	go func(c chan <- int) {
		for ; i < right && items[i] <= pivot; i++ {

		}

		log.Debugf("Next i: %d\n", i)

		c <- i;
	}(nextIChan)

	return nextIChan
}

func nextJ(j, left, pivot int, items []int) <-chan int {

	nextJChan := make(chan int)

	go func(c chan <- int) {
		for ; j > left && items[j] >= pivot; j-- {

		}

		log.Debugf("Next j: %d\n", j)

		c <- j;
	}(nextJChan)

	return nextJChan
}