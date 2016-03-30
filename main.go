package main

import (
	"github.com/buffaloluk7/parallel-quicksort/quicksort"
	"fmt"
)

func main() {
	items := []int{1,4,3,2,5}
	fmt.Printf("%v\n", items)
 	parallel_quicksort.SortQuick(items)
	fmt.Printf("%v\n", items)
}
