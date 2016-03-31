package main

import (
	"github.com/buffaloluk7/parallel-sorting-algorithms/sorting"
	"math/rand"
	"github.com/op/go-logging"
	"os"
	"time"
	"sort"
	"flag"
)

var log = logging.MustGetLogger("main")

func main() {
	// Setup logging
	var backend = logging.NewLogBackend(os.Stdout, "", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled)

	// Parse parameters
	algorithm := flag.String("algorithm", "quicksort", "The algorithm to use. Possible values: quicksort, mergesort")
	numberOfItems := *flag.Int("numberOfItems", 1000000, "The number of items to sort")
	allowDuplicates := *flag.Bool("allowDuplicates", true, "Either shuffle a range from 0 to [numberOfItems -1] or create [numberOfItems] random numbers")
	threshold := *flag.Int("threshold", 0, "The threshold to use. 0 means running all routines in parallel")
	printHelp := *flag.Bool("help", false, "Print usage")
	flag.Parse()

	if printHelp {
		flag.Usage()
		return
	}

	if numberOfItems <= 0 {
		numberOfItems = 1000000
	}
	if threshold < 0 {
		threshold = 0
	}

	// Setup items to sort
	rand.Seed(time.Now().UTC().UnixNano())
	var items []int
	if allowDuplicates {
		items = make([]int, numberOfItems)
		for i, _ := range items {
			items[i] = rand.Int() % 150
		}
	} else {
		items = rand.Perm(numberOfItems)
	}

	// Start measuring the time
	start := time.Now()

	switch *algorithm {
	case "mergesort":
		sorting.Mergesort(items, threshold)
	default:
		sorting.Quicksort(sort.IntSlice(items), threshold)
	}

	log.Infof("Took %s to sort %d items.", time.Since(start), numberOfItems)
}
