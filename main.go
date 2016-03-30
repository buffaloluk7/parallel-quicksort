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
	numberOfItems := 1000000
	items := rand.Perm(numberOfItems)

	start := time.Now()
	quicksort.Quick(sort.IntSlice(items))
	log.Infof("Took %s to sort %d items.", time.Since(start), numberOfItems)
}
