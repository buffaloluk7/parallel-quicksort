package parallel_quicksort

import (
	"sync"
	"fmt"
)

func SortQuick(items []int) {
	quicksort(0, len(items) - 1, items)
}

func quicksort(left, right int, items []int) {
	if left < right {
		pivotIndex := partition(left, right, items)
		go quicksort(left, pivotIndex - 1, items)
		go quicksort(pivotIndex + 1, right, items)
	}
}

func partition(left, right int, items []int) int {
	i := left
	j := right - 1
	pivot := items[right]
	fmt.Printf("Left: %d, right: %d, pivot value: %d (index: %d)\n",  i, j, pivot, right)

	var wg sync.WaitGroup
	wg.Add(2)

	lessThanPivot := make(chan int)
	greaterThanPivot := make(chan int)
	quit := make(chan bool)

	go func(index int, quit <- chan bool) {
		defer func() { wg.Done() }()

		for ; index < right; index++ {
			select {
			case <- quit:
				fmt.Println("i-loop received quit command\n")
				return
			default:
				fmt.Printf("Compare value %d at index %d with pivot %d\n", items[index], index, pivot)
				if items[index] > pivot {
					fmt.Printf("Value %d at index %d is greater than pivot %d\n", items[index], index, pivot)
					greaterThanPivot <- index
				}
			}
		}
		fmt.Printf("No value greater than pivot found, send index %d\n", index)
		greaterThanPivot <- index
	}(i, quit)

	go func(index int, quit <- chan bool) {
		defer func() { wg.Done() }()

		for ; index > left; index-- {
			select {
			case <- quit:
				fmt.Println("j-loop received quit command\n")
				return
			default:
				fmt.Printf("Compare value %d at index %d with pivot %d\n", items[index], index, pivot)
				if items[j] < pivot {
					fmt.Printf("Value %d at index %d is smaller than pivot %d\n", items[index], index, pivot)
					lessThanPivot <- index
				}
			}
		}
		fmt.Printf("No value smaller than pivot found, send index %d\n", index)
		lessThanPivot <- index
	}(j, quit)

	for {
		select {
		case lessThanPivotIndex := <-lessThanPivot:
			greaterThanPivotIndex := <- greaterThanPivot
			fmt.Printf("Got two values indicies: (greater) %d, (less) %d\n", greaterThanPivotIndex, lessThanPivotIndex)
			if greaterThanPivotIndex < lessThanPivotIndex {
				fmt.Printf("Swap values\n")
				items[greaterThanPivotIndex], items[lessThanPivotIndex] = items[lessThanPivotIndex], items[greaterThanPivotIndex]
			} else {
				fmt.Printf("j index is greater than i index -> stop for loop\n")
				//quit <- true
				//quit <- true
				close(quit)
				fmt.Printf("Close quit channel.\n")
				lastI := greaterThanPivotIndex

				//wg.Wait()

				if items[lastI] > pivot {
					items[lastI], items[right] = items[right], items[lastI]
				}

				return lastI
			}

		case greaterThanPivotIndex := <- greaterThanPivot:
			lessThanPivotIndex := <- lessThanPivot
			fmt.Printf("Got two values indicies: (greater) %d, (less) %d\n", greaterThanPivotIndex, lessThanPivotIndex)
			if greaterThanPivotIndex < lessThanPivotIndex {
				fmt.Printf("Swap values\n")
				items[greaterThanPivotIndex], items[lessThanPivotIndex] = items[lessThanPivotIndex], items[greaterThanPivotIndex]
			} else {
				fmt.Printf("j index is greater than i index -> stop for loop\n")
				//quit <- true
				//quit <- true
				close(quit)
				fmt.Printf("Close quit channel.\n")
				lastI := greaterThanPivotIndex

				//wg.Wait()

				if items[lastI] > pivot {
					items[lastI], items[right] = items[right], items[lastI]
				}

				return lastI
			}
		}
	}
}