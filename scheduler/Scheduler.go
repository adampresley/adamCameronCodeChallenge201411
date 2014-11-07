package scheduler

import (
	"sync"

	"github.com/adampresley/adamCameronCodeChallenge201411/subarray"
)

/*
This function takes an array of integers and iterates over them.
Each iteration starts at an index one greater than the previous
index, createing a goroutine to sum up those number slices. The final
result is an array of SubArrayResult structures.
*/
func GetSliceCounts(numberArray []int, threshold int) []subarray.SubArrayResult {
	var wg sync.WaitGroup

	resultChannel := make(chan subarray.SubArrayResult, 100)
	doneChannel := make(chan bool)

	result := make([]subarray.SubArrayResult, 0)

	/*
	 * Listen on this results channel for totals
	 * and add them to our final results array.
	 */
	go func() {
		var item subarray.SubArrayResult

		for {
			select {
			case item = <- resultChannel:
				result = append(result, item)

			case <- doneChannel:
				break
			}
		}
	}()

	for index, _ := range numberArray {
		wg.Add(1)

		go subarray.SumArray(numberArray[index:], threshold, resultChannel, &wg)
	}

	wg.Wait()
	doneChannel <- true

	return result
}
