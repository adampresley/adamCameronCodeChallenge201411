package subarray

type SubArrayResult struct {
	Sequence []int
	Total    int
}

func NewSubArrayResult(sequence []int) SubArrayResult {
	total := 0

	for _, value := range sequence {
		total += value
	}

	return SubArrayResult{
		Sequence: sequence,
		Total:    total,
	}
}