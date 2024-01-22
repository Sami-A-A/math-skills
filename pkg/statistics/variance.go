package statistics

func Variance(numSet []float64) float64 {

	//Formula for variance: (Σ((x - μ)²)) / n
	//1. Find the mean of the data set 									-> μ
	//2. Subtract each number (x) by the mean (μ) 						-> x - μ
	//3. Get the deviation scores by squaring the results from step 2	-> (x - μ)²
	//4. Find the average of the deviation scores						-> (Σ((x - μ)²)) / n

	var variance float64
	var deviationScores []float64

	mean := Average(numSet)
	
	for _, n := range numSet{
		n -= mean
		deviationScores = append(deviationScores, n * n)
	}

	variance = Average(deviationScores)
	
	return variance
}