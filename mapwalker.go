package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aybabtme/uniplot/histogram"
)

func mapIterate(initialSize int, numIterations int) []float64 {
	results := make([]float64, numIterations)
	for iteration := 0; iteration < numIterations; iteration++ {
		// initialize
		m := make(map[int]int)
		for i := 0; i < initialSize; i++ {
			m[i] = i
		}

		// iterate the map, adding a value each time
		for k := range m {
			m[k*2] = k * 2
		}
		results[iteration] = float64(len(m))
	}
	return results
}

type result struct {
	initialSize   int
	mapFinalSizes []float64 // The histogram library wants floats
}

func usage() {
	fmt.Print(`
mapwalker is a program to answer a fun toy question: what happens as you insert
into a golang map while traversing it? It gives a histogram of final sizes from
a real golang map, and a simulated map.

`)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	initialSize := flag.Int("initial", 1024, "initial size of map")
	numIterations := flag.Int("iterations", 1000, "number of iterations to run")
	flag.Parse()

	mapFinalSizes := mapIterate(*initialSize, *numIterations)

	fmt.Printf("Initial size of map: %d\n", *initialSize)
	fmt.Printf("Distribution of final size of map over %d iterations\n", *numIterations)
	hist := histogram.Hist(10, mapFinalSizes)
	histogram.Fprint(os.Stdout, hist, histogram.Linear(20))
}
