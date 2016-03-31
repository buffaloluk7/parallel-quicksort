package sorting

import (
	"sync"
)

func Mergesort(list []int, threshold int) []int {
	size := len(list)
	if size <= 1 {
		return list
	}

	middle := size / 2
	var left, right []int

	sortParallel := threshold == 0 || size > threshold
	if sortParallel {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			left = Mergesort(list[:middle], threshold)

		}()

		go func() {
			defer wg.Done()
			right = Mergesort(list[middle:], threshold)
		}()

		wg.Wait()
	} else {
		left = Mergesort(list[:middle], threshold)
		right = Mergesort(list[middle:], threshold)
	}

	return merge(left, right)
}

func merge(leftList, rightList []int) []int {
	size := len(leftList) + len(rightList)
	slice := make([]int, size)
	i, j := 0, 0

	for k := 0; k < size; k++ {
		if i > len(leftList) - 1 && j <= len(rightList) - 1 {
			slice[k] = rightList[j]
			j++
		} else if j > len(rightList) - 1 && i <= len(leftList) - 1 {
			slice[k] = leftList[i]
			i++
		} else if leftList[i] < rightList[j] {
			slice[k] = leftList[i]
			i++
		} else {
			slice[k] = rightList[j]
			j++
		}
	}

	return slice
}