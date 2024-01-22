package statistics

func Variance(numSet []float64) float64 {
	var variance float64
	var deviationScores []float64
	var sumOfDeviationScores float64

	mean := Average(numSet)
	
	for _, n := range numSet{
		deviationScores = append(deviationScores, n - mean)
	}

	for _, n := range deviationScores {
		sumOfDeviationScores += n
	}

	variance = (sumOfDeviationScores * sumOfDeviationScores) / float64(len(numSet))
	
	return variance
}