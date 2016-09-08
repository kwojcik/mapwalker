package main

import (
	"flag"
	"fmt"
	"math"
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
			m[k+initialSize] = k
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

func main() {
	flag.Usage = usage
	initialSize := flag.Int("initial", 1024, "initial size of map")
	numIterations := flag.Int("iterations", 1000, "number of iterations to run")
	flag.Parse()

	mapFinalSizes := mapIterate(*initialSize, *numIterations)
	stddev, avg := maths(mapFinalSizes)
	fmt.Printf("Initial size of map: %d\n", *initialSize)
	fmt.Printf("Iterations: %d\n", *numIterations)
	fmt.Printf("Final size:\n")
	fmt.Printf("\tAverage\t%f\n", avg)
	fmt.Printf("\tStddev\t%f\n", stddev)
	fmt.Printf("\nDistribution of final size of map\n")
	hist := histogram.Hist(10, mapFinalSizes)
	histogram.Fprint(os.Stdout, hist, histogram.Linear(20))
}
