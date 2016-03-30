package main

import (
	"github.com/buffaloluk7/parallel-quicksort/quicksort"
	"math/rand"
	"github.com/op/go-logging"
	"os"
	"time"
)

var log = logging.MustGetLogger("main")

func main() {
	var backend = logging.NewLogBackend(os.Stdout, "", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled)

	numberOfItems := 1000000
	items := rand.Perm(numberOfItems)

	start := time.Now()
 	parallel_quicksort.SortQuick(items, false)
	log.Infof("Took %s to sort %d items.", time.Since(start), numberOfItems)
}
