package subarray

import (
	"sync"
)

/*
This function takes an array of SubArrayResult structures and returns the
one structure that contains the longest sequence of numbers. If there is
more than one sequence that has the same lenght, the one with the biggest
total wins.
*/
func GetLongestSequence(subArrayResults []SubArrayResult) SubArrayResult {
	longestIndex := 0
	longestLength := 0

	for index, item := range subArrayResults {
		if len(item.Sequence) > longestLength {
			longestLength = len(item.Sequence)
			longestIndex = index
		}

		if len(item.Sequence) == longestLength && item.Total > subArrayResults[longestIndex].Total {
			longestLength = len(item.Sequence)
			longestIndex = index
		}
	}

	return subArrayResults[longestIndex]
}

/*
This function takes an array of integers and a threshold and sums the integers
up, making sure not to exceed the threshold. The numbers that were summed
and total sum are captured in a SubArraySequence structure and written to
a channel.
*/
func SumArray(numberArray []int, threshold int, resultChannel chan SubArrayResult, wg *sync.WaitGroup) {
	total := 0
	sequence := make([]int, 0)

	defer wg.Done()

	for _, value := range numberArray {
		if (total + value) <= threshold {
			total += value
			sequence = append(sequence, value)
		} else {
			break
		}
	}

	if len(sequence) > 0 {
		result := NewSubArrayResult(sequence)
		resultChannel <- result
	}
}
