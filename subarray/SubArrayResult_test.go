package subarray

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubArrayResult(t *testing.T) {
	Convey("Calling NewSubArrayResult should return a new struct with a valid total...", t, func() {
		sequence := []int{1, 2, 3, 4, 5,}

		expected := SubArrayResult{
			Sequence: sequence,
			Total:    15,
		}

		actual := NewSubArrayResult(sequence)
		So(actual, ShouldResemble, expected)
	})
}
