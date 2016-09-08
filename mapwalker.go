package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/kwojcik/mapwalker/internal/nogrowmap"

	"github.com/aybabtme/uniplot/histogram"
)

func mapRun(initialSize int) int {
	// initialize
	m := make(map[int]int)
	for i := 0; i < initialSize; i++ {
		m[i] = i
	}

	// iterate the map, adding a value each time
	for k := range m {
		m[k+initialSize] = k
	}
	return len(m)
}

func noGrowMapRun(initialSize int) int {
	m := nogrowmap.NewNoGrowMap(initialSize)
	for range m.Iterator {
		m.Insert()
	}
	return m.Size
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

// maths does some maths
func maths(nums []float64) (stddev float64, avg float64) {
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	avg = sum / float64(len(nums))

	sumOfDiffsSquared := 0.0
	for _, n := range nums {
		sumOfDiffsSquared += math.Pow(float64(n)-avg, 2.0)
	}

	// or is it len(nums) - 1 ...?
	stddev = math.Sqrt(sumOfDiffsSquared / float64(len(nums)))
	return stddev, avg
}

func printResults(mapType string, initialSize int, numIterations int, results []float64) {
	stddev, avg := maths(results)
	fmt.Printf("\nResults for %s\n", mapType)
	fmt.Printf("Initial size: %d\n", initialSize)
	fmt.Printf("Iterations: %d\n", numIterations)
	fmt.Printf("Final size:\n")
	fmt.Printf("\tAverage\t%f\n", avg)
	fmt.Printf("\tStddev\t%f\n", stddev)
	fmt.Printf("Distribution of final size of map\n")
	hist := histogram.Hist(15, results)
	histogram.Fprint(os.Stdout, hist, histogram.Linear(20))
}

func main() {
	flag.Usage = usage
	initialSize := flag.Int("initial", 1024, "initial size of map")
	numIterations := flag.Int("iterations", 1000, "number of iterations to run")
	flag.Parse()

	mapResults := make([]float64, *numIterations)
	noGrowMapResults := make([]float64, *numIterations)
	for i := 0; i < *numIterations; i++ {
		mapResults[i] = float64(mapRun(*initialSize))
		noGrowMapResults[i] = float64(noGrowMapRun(*initialSize))
	}

	printResults("map", *initialSize, *numIterations, mapResults)
	printResults("NoGrowMap", *initialSize, *numIterations, noGrowMapResults)
}
