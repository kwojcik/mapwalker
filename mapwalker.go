package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/kwojcik/mapwalker/internal/nogrowmap"

	"github.com/aybabtme/uniplot/histogram"
)

func mapRun(initialSize int, capacity int) int {
	// initialize
	var m map[int]int
	if capacity == 0 {
		m = make(map[int]int)
	} else {
		m = make(map[int]int, capacity)
	}
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

func printResults(mapType string, initialSize int, numIterations int, results []float64,
	rawResults bool) {
	if rawResults {
		fmt.Println(strings.Trim(fmt.Sprint(results), "[]"))
	} else {
		stddev, avg := maths(results)
		fmt.Printf("\nResults for %s\n", mapType)
		fmt.Printf("Initial size: %d\n", initialSize)
		fmt.Printf("Iterations: %d\n", numIterations)
		fmt.Printf("Final size:\n")
		fmt.Printf("\tAverage\t%f\n", avg)
		fmt.Printf("\tStddev\t%f\n", stddev)
		fmt.Printf("Distribution of final size of map\n")
		hist := histogram.Hist(10, results)
		histogram.Fprint(os.Stdout, hist, histogram.Linear(20))
	}
}

func runIterations(numIterations int, runFunc func() int) []float64 {
	results := make([]float64, numIterations)
	for i := 0; i < numIterations; i++ {
		results[i] = float64(runFunc())
	}
	return results
}

func main() {
	flag.Usage = usage
	initialSize := flag.Int("initial", 1024, "initial size of map")
	numIterations := flag.Int("iterations", 1000, "number of iterations to run")
	onlyMap := flag.Bool("onlyMap", false, "only run for a normal map")
	onlyNoGrowMap := flag.Bool("onlyNoGrowMap", false, "only run for a simulated no-grow map")
	onlySparseMap := flag.Bool("onlySparseMap", false, "only run for a sparse map")
	rawResults := flag.Bool("rawResults", false, "only show raw final sizes")
	flag.Parse()

	all := !*onlyMap && !*onlyNoGrowMap && !*onlySparseMap
	if all || *onlyMap {
		mapResults := runIterations(*numIterations,
			func() int {
				return mapRun(*initialSize, *initialSize)
			})
		printResults("map", *initialSize, *numIterations, mapResults, *rawResults)
	}
	if all || *onlyNoGrowMap {
		noGrowMapResults := runIterations(*numIterations,
			func() int {
				return noGrowMapRun(*initialSize)
			})
		printResults("NoGrowMap", *initialSize, *numIterations, noGrowMapResults, *rawResults)
	}
	if all || *onlySparseMap {
		sparseMapResults := runIterations(*numIterations,
			func() int {
				return mapRun(*initialSize, *initialSize*50)
			})
		printResults("sparse map", *initialSize, *numIterations, sparseMapResults, *rawResults)
	}
}
