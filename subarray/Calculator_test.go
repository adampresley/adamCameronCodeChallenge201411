package subarray

import (
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculator(t *testing.T) {
	Convey("GetLongestSequence()...", t, func() {
		Convey("Returns longest array sequence", func() {
			// 1, 2, 3, 4, 5
			// Threshold 10

			expected := SubArrayResult{
				Sequence: []int{1, 2, 3, 4,},
				Total:    10,
			}

			subArrayResults := []SubArrayResult{
				SubArrayResult{Sequence: []int{1, 2, 3, 4,}, Total: 10,},
				SubArrayResult{Sequence: []int{2, 3, 4,}, Total: 9,},
				SubArrayResult{Sequence: []int{3, 4,}, Total: 7,},
				SubArrayResult{Sequence: []int{4, 5,}, Total: 9,},
			}

			actual := GetLongestSequence(subArrayResults)
			So(actual, ShouldResemble, expected)
		})

		Convey("Returns greater total when more than one array sequences are of same length", func() {
			// 10, 10, 11, 10, 10, 12
			// Threshold 32

			expected := SubArrayResult{
				Sequence: []int{10, 10, 12,},
				Total:    32,
			}

			subArrayResults := []SubArrayResult{
				SubArrayResult{Sequence: []int{10, 10, 11, }, Total: 31,},
				SubArrayResult{Sequence: []int{10, 11, 10,}, Total: 31,},
				SubArrayResult{Sequence: []int{11, 10, 10,}, Total: 31,},
				SubArrayResult{Sequence: []int{10, 10, 12,}, Total: 32,},
			}

			actual := GetLongestSequence(subArrayResults)
			So(actual, ShouldResemble, expected)
		})
	})

	Convey("SumArray()...", t, func() {
		Convey("Will sum values in an integer array up until the threshold", func() {
			sequence := []int{1, 2, 3, 4, 5,}
			threshold := 6
			channel := make(chan SubArrayResult, 10)

			/*
			 * We need a wait group cause SumArray() alerts a wait group
			 * when it is done.
			 */
			var waitGroup sync.WaitGroup
			waitGroup.Add(1)

			expected := SubArrayResult{Sequence: []int{1, 2, 3,}, Total: 6,}
			SumArray(sequence, threshold, channel, &waitGroup)

			waitGroup.Wait()

			actual := <- channel
			So(actual, ShouldResemble, expected)
		})
	})
}
