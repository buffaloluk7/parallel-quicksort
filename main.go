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
	numberOfItems := 1000000
	items := rand.Perm(numberOfItems)
	items2 := rand.Perm(numberOfItems)

	start2 := time.Now()
	parallel_quicksort.SortQuick(items, true)
	log.Infof("Took %s to sort %d items.", time.Since(start2), numberOfItems)

	start := time.Now()
	sorting.Quicksort(sort.IntSlice(items2))
	log.Infof("Took %s to sort %d items.", time.Since(start), numberOfItems)
}
