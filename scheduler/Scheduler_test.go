package scheduler

import (
	"testing"

	"github.com/adampresley/adamCameronCodeChallenge201411/subarray"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSliceCounts(t *testing.T) {
	Convey("GetSliceCounts will return an array of SubArrayResult structures, each with a slice of arrays with totals not exceeding threshold", t, func() {
		// 10, 10, 11, 10, 10, 12
		// Threshold 32

		sequence := []int{10, 10, 11, 10, 10, 12,}
		threshold := 32

		expected := []subarray.SubArrayResult{
			subarray.SubArrayResult{Sequence: []int{10, 10, 11, }, Total: 31,},
			subarray.SubArrayResult{Sequence: []int{10, 11, 10,}, Total: 31,},
			subarray.SubArrayResult{Sequence: []int{11, 10, 10,}, Total: 31,},
			subarray.SubArrayResult{Sequence: []int{10, 10, 12,}, Total: 32,},
			subarray.SubArrayResult{Sequence: []int{10, 12,}, Total: 22,},
			subarray.SubArrayResult{Sequence: []int{12,}, Total: 12,},
		}

		actual := GetSliceCounts(sequence, threshold)
		So(actual, ShouldResemble, expected)
	})
}