package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adampresley/adamCameronCodeChallenge201411/scheduler"
	"github.com/adampresley/adamCameronCodeChallenge201411/subarray"
)

var input = flag.String("input", "", "Comma delimited list of integer numbers")
var threshold = flag.Int("threshold", 10, "Integer number for array slice to not exceed")

func main() {
	flag.Parse()

	/*
	 * Array of integers to determine longest sequence from. These are a string
	 * list on the command line that are converted to integers.
	 */
	inputArray := make([]int, 0)

	for _, value := range strings.Split(*input, ",") {
		i, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("ERROR - Unable to convert", value, "to integer")
			os.Exit(1)
		}

		inputArray = append(inputArray, i)
	}

	/*
	 * Use the GetSliceCounts() method to get all possible
	 * sequence slices. This is done using goroutines (in parallel).
	 * The result is then passed to GetLongestSequence to figure
	 * out which one wins.
	 */
	subseriesResults := scheduler.GetSliceCounts(inputArray, *threshold)
	subseries := subarray.GetLongestSequence(subseriesResults)

	fmt.Println("")
	fmt.Println("Sequence", subseries.Sequence, "wins with a length of", len(subseries.Sequence), " for a total of", subseries.Total)
}