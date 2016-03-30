package main

import (
	"math/rand"
	"github.com/op/go-logging"
	"os"
	"time"
	"github.com/Lacrymology/golang-examples/sorting"
	"sort"
	"github.com/buffaloluk7/parallel-sorting-algorithms/quicksort"
)

var log = logging.MustGetLogger("main")

func main() {
	var backend = logging.NewLogBackend(os.Stdout, "", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled)

	// see: https://github.com/Lacrymology/golang-examples/blob/master/sorting/example/test.go
	rand.Seed(time.Now().UTC().UnixNano())
	numberOfItems := 10000
	items := make([]int, numberOfItems)
	for i, _ := range items {
		items[i] = rand.Int() % 200
	}

	items2 := make([]int, numberOfItems)
	copy(items2, items)
	
	start := time.Now()
	sorting.Quicksort(sort.IntSlice(items2))
	log.Infof("Took %s to sort %d items.", time.Since(start), numberOfItems)

	start2 := time.Now()
	quicksort.Quick(sort.IntSlice(items))
	log.Infof("Took %s to sort %d items.", time.Since(start2), numberOfItems)
}
