package main

import (
	"flag"
	"fmt"
	"github.com/moritzschramm/sorting-algos-go/algos"
	"github.com/moritzschramm/sorting-algos-go/generator"
	"os"
	"time"
)

var (
	LENGTH   = int(1e2)
	MAX_RAND = int(1e3)
)

func main() {

	algoName := parseFlags()

	a := generator.Generate(LENGTH, MAX_RAND)
	algo := getSortingAlgorithmByName(algoName)

	fmt.Println(a)
	fmt.Println("---------------------------------------------------")

	start := time.Now()
	algo(a)
	elapsed := time.Since(start)

	fmt.Println(a)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func listAlgos() {

	fmt.Println("insertion-sort")
	fmt.Println("merge-sort")
	fmt.Println("quicksort")
	fmt.Println("heapsort")
	//fmt.Println("timsort")
}

func getSortingAlgorithmByName(name string) func([]int) {

	switch name {

	case "insertion-sort":
		return algos.InsertionSort
	case "merge-sort":
		return algos.MergeSort
	case "quicksort":
		return algos.QuickSort
	case "heapsort":
		return algos.HeapSort
	//case "timsort":
	//	return algos.TimSort
	default:
		fmt.Println("Unknown algorithm")
		os.Exit(1)
		return func(a []int) {}
	}
}

func parseFlags() string {

	length := flag.Int("len", LENGTH, "The length of the array to be sorted")
	max := flag.Int("rand", MAX_RAND, "The maximum randomly generated value")

	flag.Parse()

	LENGTH = *length
	MAX_RAND = *max

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Expecting _exactly_ one sorting algorithm as argument. Use argument 'list' to list all algorithms.")
		os.Exit(1)
	}
	name := args[0]
	if name == "list" {
		listAlgos()
		os.Exit(0)
	}

	return name
}
