package statistics

import (
	"sort"
)

func Median(numSet []float64) float64 {
	var median float64

	// Sorting numSet
	sort.Float64s(numSet)

	// Even number set
	if len(numSet) % 2 == 0{
		medianA := numSet[len(numSet)/2]
		medianB := numSet[len(numSet)/2 - 1]
		median = (medianA + medianB)/2
	}

	// Odd number set
	if len(numSet) % 2 != 0 {
		median = numSet[len(numSet)/2]
	}
	return median
}